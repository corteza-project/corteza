package rbac

import (
	"context"
	"sync"
	"time"

	"github.com/cortezaproject/corteza-server/pkg/sentry"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	service struct {
		l      *sync.Mutex
		logger *zap.Logger

		//  service will flush values on TRUE or just reload on FALSE
		f chan bool

		rules RuleSet

		store rbacRulesStore
	}

	// RuleFilter is a dummy struct to satisfy store codegen
	RuleFilter struct{}

	Controller interface {
		Can(roles []uint64, res Resource, op Operation, ff ...CheckAccessFunc) bool
		Check(res Resource, op Operation, roles ...uint64) (v Access)
		Grant(ctx context.Context, wl Whitelist, rules ...*Rule) (err error)
		Watch(ctx context.Context)
		FindRulesByRoleID(roleID uint64) (rr RuleSet)
		SignificantRoles(res Resource, op Operation) (aRR, dRR []uint64)
		Rules() (rr RuleSet)
		Reload(ctx context.Context)
	}
)

var (
	// Global RBAC service
	gRBAC Controller
)

const (
	watchInterval = time.Hour
)

// Global returns global RBAC service
func Global() Controller {
	return gRBAC
}

func SetGlobal(svc Controller) {
	gRBAC = svc
}

func Initialize(logger *zap.Logger, s rbacRulesStore) error {
	if gRBAC != nil {
		// Prevent multiple initializations
		return nil
	}

	SetGlobal(NewService(logger, s))
	return nil
}

// NewService initializes service{} struct
//
// service{} struct preloads, checks, grants and flushes privileges to and from store
// It acts as a caching layer
func NewService(logger *zap.Logger, s rbacRulesStore) (svc *service) {
	svc = &service{
		l: &sync.Mutex{},
		f: make(chan bool),

		store: s,
	}

	if logger != nil {
		svc.logger = logger.Named("rbac")
	}

	return
}

// Can function performs permission check for roles in context
//
// First extracts roles from context, then
// use Check() to test against permission rules and
// iterate over all fallback functions
//
// System user is always allowed to do everything
//
// When not explicitly allowed through rules or fallbacks, function will return FALSE.
func (svc service) Can(roles []uint64, res Resource, op Operation, ff ...CheckAccessFunc) bool {
	// Checking rules
	var v = svc.Check(res.RBACResource(), op, roles...)
	if v != Inherit {
		return v == Allow
	}

	// Checking fallback functions
	for _, f := range ff {
		v = f()

		if v != Inherit {
			return v == Allow
		}
	}

	return false
}

// Check verifies if role has access to perform an operation on a resource
//
// See RuleSet's Check() func for details
func (svc service) Check(res Resource, op Operation, roles ...uint64) (v Access) {
	svc.l.Lock()
	defer svc.l.Unlock()

	return svc.rules.Check(res, op, roles...)
}

// Grant appends and/or overwrites internal rules slice
//
// All rules with Inherit are removed
func (svc *service) Grant(ctx context.Context, wl Whitelist, rules ...*Rule) (err error) {
	svc.l.Lock()
	defer svc.l.Unlock()

	if err = svc.checkRules(wl, rules...); err != nil {
		return err
	}

	svc.grant(rules...)

	return svc.flush(ctx)
}

func (svc service) checkRules(wl Whitelist, rules ...*Rule) error {
	for _, r := range rules {
		if !wl.Check(r) {
			return errors.Errorf("invalid rule: '%s' on '%s'", r.Operation, r.Resource)
		}
	}

	return nil
}

func (svc *service) grant(rules ...*Rule) {
	svc.rules = svc.rules.Merge(rules...)
}

// Watches for changes
func (svc service) Watch(ctx context.Context) {
	go func() {
		defer sentry.Recover()

		var ticker = time.NewTicker(watchInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				svc.Reload(ctx)
			case <-svc.f:
				svc.Reload(ctx)
			}
		}
	}()

	svc.logger.Debug("watcher initialized")
}

func (svc service) FindRulesByRoleID(roleID uint64) (rr RuleSet) {
	svc.l.Lock()
	defer svc.l.Unlock()

	rr, _ = svc.rules.Filter(func(rule *Rule) (b bool, e error) {
		return rule.RoleID == roleID, nil
	})

	return
}

func (svc service) Rules() (rr RuleSet) {
	svc.l.Lock()
	defer svc.l.Unlock()
	return svc.rules
}

func (svc *service) Reload(ctx context.Context) {
	svc.l.Lock()
	defer svc.l.Unlock()

	rr, _, err := svc.store.SearchRbacRules(ctx, RuleFilter{})
	svc.logger.Debug(
		"reloading rules",
		zap.Error(err),
		zap.Int("before", len(svc.rules)),
		zap.Int("after", len(rr)),
	)

	if err == nil {
		svc.rules = rr
	}
}

// SignificantRoles returns two list of significant roles.
//
// See sigRoles on rules for more details
func (svc *service) SignificantRoles(res Resource, op Operation) (aRR, dRR []uint64) {
	svc.l.Lock()
	defer svc.l.Unlock()

	return svc.rules.sigRoles(res, op)
}

func (svc service) flush(ctx context.Context) (err error) {
	d, u := svc.rules.Dirty()

	err = svc.store.DeleteRbacRule(ctx, d...)
	if err != nil {
		return
	}

	err = svc.store.UpsertRbacRule(ctx, u...)
	if err != nil {
		return
	}

	u.Clear()
	svc.rules = u
	svc.logger.Debug("flushed rules",
		zap.Int("updated", len(u)),
		zap.Int("deleted", len(d)))

	return
}
