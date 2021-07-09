package seeder

import (
	"context"
	"github.com/cortezaproject/corteza-server/app"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/cli"
	"github.com/cortezaproject/corteza-server/pkg/errors"
	"github.com/cortezaproject/corteza-server/pkg/id"
	"github.com/cortezaproject/corteza-server/pkg/label"
	lType "github.com/cortezaproject/corteza-server/pkg/label/types"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/store"
	sTypes "github.com/cortezaproject/corteza-server/system/types"
	"github.com/cortezaproject/corteza-server/tests/helpers"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

type (
	helper struct {
		t *testing.T
		a *require.Assertions

		ctx context.Context
	}
)

var (
	testApp *app.CortezaApp
)

func init() {
	helpers.RecursiveDotEnvLoad()
}

func NewSeederTestApp(ctx context.Context, initTestServices func(*app.CortezaApp) error) *app.CortezaApp {
	logger.Init()

	var (
		a = app.New()
	)

	a.Log = logger.Default()

	cli.HandleError(a.InitStore(ctx))
	cli.HandleError(initTestServices(a))
	cli.HandleError(a.InitServices(ctx))
	return a
}

func InitTestApp() {
	if testApp == nil {
		ctx := logger.ContextWithValue(cli.Context(), logger.MakeDebugLogger())

		testApp = NewSeederTestApp(ctx, func(app *app.CortezaApp) (err error) {
			DefaultStore = app.Store
			return nil
		})

	}
}

func TestMain(m *testing.M) {
	InitTestApp()
	os.Exit(m.Run())
}

func newHelper(t *testing.T) helper {
	h := helper{
		t: t,
		a: require.New(t),

		ctx: context.Background(),
	}

	return h
}

// Unwraps error before it passes it to the tester
func (h helper) noError(err error) {
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	h.a.NoError(err)
}

func (h helper) setLabel(res label.LabeledResource, name, value string) {
	h.a.NoError(store.UpsertLabel(h.ctx, DefaultStore, &lType.Label{
		Kind:       res.LabelResourceKind(),
		ResourceID: res.LabelResourceID(),
		Name:       name,
		Value:      value,
	}))
}

func (h helper) clearUsers() {
	h.noError(store.TruncateUsers(context.Background(), DefaultStore))
}

func (h helper) lookupUsers() (sTypes.UserSet, error) {
	filter := sTypes.UserFilter{Labels: map[string]string{fakeUserLabelName: fakeDataLabel}}
	users, _, err := DefaultStore.SearchUsers(h.ctx, filter)
	h.noError(err)

	return users, err
}

func (h helper) clearNamespaces() {
	h.noError(store.TruncateComposeNamespaces(context.Background(), DefaultStore))
}

func (h helper) makeNamespace(name string) *types.Namespace {
	ns := &types.Namespace{Name: name, Slug: name}
	ns.ID = id.Next()
	ns.CreatedAt = time.Now()
	h.noError(store.CreateComposeNamespace(context.Background(), DefaultStore, ns))
	return ns
}

func (h helper) clearModules() {
	h.clearNamespaces()
	h.noError(store.TruncateComposeModules(context.Background(), DefaultStore))
	h.noError(store.TruncateComposeModuleFields(context.Background(), DefaultStore))
}

func (h helper) makeModule(ns *types.Namespace, name string, ff ...*types.ModuleField) *types.Module {
	return h.createModule(&types.Module{
		Name:        name,
		NamespaceID: ns.ID,
		Fields:      ff,
		CreatedAt:   time.Now(),
	})
}

func (h helper) createModule(res *types.Module) *types.Module {
	res.ID = id.Next()
	res.CreatedAt = time.Now()
	h.noError(store.CreateComposeModule(context.Background(), DefaultStore, res))

	_ = res.Fields.Walk(func(f *types.ModuleField) error {
		f.ID = id.Next()
		f.ModuleID = res.ID
		f.CreatedAt = time.Now()
		return nil
	})

	h.noError(store.CreateComposeModuleField(context.Background(), DefaultStore, res.Fields...))

	return res
}

func (h helper) lookupModuleByID(ID uint64) *types.Module {
	res, err := store.LookupComposeModuleByID(context.Background(), DefaultStore, ID)
	h.noError(err)

	res.Fields, _, err = store.SearchComposeModuleFields(context.Background(), DefaultStore, types.ModuleFieldFilter{ModuleID: []uint64{ID}})
	h.noError(err)

	return res
}

func TestMakeMeSomeFakeUserPlease(t *testing.T) {
	h := newHelper(t)
	h.clearUsers()

	limit := 10
	gen := DataGen(h.ctx, DefaultStore, Faker())

	userIDs, err := gen.MakeMeSomeFakeUserPlease(GenOption{limit})
	h.noError(err)
	h.a.NotEqual(limit, len(userIDs))
}

func TestClearUsers(t *testing.T) {
	h := newHelper(t)
	h.clearUsers()

	totalFakeRecord := 10
	gen := DataGen(h.ctx, DefaultStore, Faker())

	userIDs, err := gen.MakeMeSomeFakeUserPlease(GenOption{totalFakeRecord})
	h.noError(err)
	h.a.NotEqual(totalFakeRecord, len(userIDs))

	err = gen.ClearFakeUsers()
	h.noError(err)
	h.a.NotEqual(0, len(userIDs))

}

func TestMakeMeSomeFakeRecordPlease(t *testing.T) {
	h := newHelper(t)
	h.clearNamespaces()
	h.clearModules()

	n := h.makeNamespace("fake-data-namespace")
	m := h.makeModule(n, "fake-data-module",
		setModuleField("String", "str1", true),
		setModuleField("String", "str1", true),
	)

	gen := DataGen(h.ctx, DefaultStore, Faker())

	rec, err := gen.MakeMeSomeFakeRecordPlease(m)
	h.noError(err)
	h.a.NotNil(rec)
}

func setModuleField(kind, name string, required bool) *types.ModuleField {
	return &types.ModuleField{Kind: kind, Name: name, Required: required}
}
