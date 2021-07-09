package seeder

import (
	"context"
	"fmt"
	cTypes "github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/id"
	lTypes "github.com/cortezaproject/corteza-server/pkg/label/types"
	"github.com/cortezaproject/corteza-server/store"
	sTypes "github.com/cortezaproject/corteza-server/system/types"
	"time"
)

type (
	seeder struct {
		ctx   context.Context
		store storeService
		faker fakerService
	}
	Params struct {
		// (optional) no record to be generate; Default value will be 1
		Limit int
	}

	fakerService interface {
		fakeValueByName(name string) (val string, ok bool)
		fakeValue(name, kind string, opt valueOptions) (val string, err error)
		fakeUserHandle(s string) string
	}

	storeService interface {
		UpsertLabel(ctx context.Context, rr ...*lTypes.Label) error
		SearchUsers(ctx context.Context, f sTypes.UserFilter) (sTypes.UserSet, sTypes.UserFilter, error)
		CreateUser(ctx context.Context, rr ...*sTypes.User) error
		DeleteUser(ctx context.Context, rr ...*sTypes.User) error
		LookupComposeModuleByID(ctx context.Context, id uint64) (*cTypes.Module, error)
		SearchComposeRecords(ctx context.Context, _mod *cTypes.Module, f cTypes.RecordFilter) (cTypes.RecordSet, cTypes.RecordFilter, error)
		CreateComposeRecord(ctx context.Context, mod *cTypes.Module, rr ...*cTypes.Record) error
		DeleteComposeRecord(ctx context.Context, m *cTypes.Module, rr ...*cTypes.Record) error
	}
)

var (
	DefaultStore store.Storer
)

const (
	fakeDataLabel = "generated"
)

func Seeder(ctx context.Context, store store.Storer, faker fakerService) *seeder {
	DefaultStore = store
	return &seeder{ctx, store, faker}
}

// getLimit return data generation limit; It will return Default(1) if limit is 0
func (p Params) getLimit() int {
	if p.Limit == 0 {
		return 1
	}
	return p.Limit
}

// CreateLabel return the label for generate data
func (s seeder) CreateLabel(resourceID uint64, kind, name string) *lTypes.Label {
	return &lTypes.Label{
		Kind:       kind,
		ResourceID: resourceID,
		Name:       name,
		Value:      fakeDataLabel,
	}
}

// CreateUser creates given no of users into DB
func (s seeder) CreateUser(params Params) (IDs []uint64, err error) {
	var users []*sTypes.User
	var labels []*lTypes.Label

	for i := 0; i < params.getLimit(); i++ {
		var user sTypes.User
		user.ID = id.Next()
		user.Email, _ = s.faker.fakeValueByName("Email")
		user.Name, _ = s.faker.fakeValueByName("Name")
		user.Handle = s.faker.fakeUserHandle(user.Name)
		user.Kind = sTypes.NormalUser
		user.CreatedAt = time.Now()

		IDs = append(IDs, user.ID)
		users = append(users, &user)
		labels = append(labels, s.CreateLabel(
			user.ID,
			user.LabelResourceKind(),
			fakeDataLabel,
		))
	}

	err = s.store.CreateUser(s.ctx, users...)
	if err != nil {
		return
	}

	err = s.store.UpsertLabel(s.ctx, labels...)
	if err != nil {
		return
	}

	return
}

// DeleteAllUser deletes all the fake users from DB
func (s seeder) DeleteAllUser() (err error) {
	filter := sTypes.UserFilter{
		Labels: map[string]string{
			fakeDataLabel: fakeDataLabel,
		},
	}
	users, _, err := s.store.SearchUsers(s.ctx, filter)
	if err != nil {
		return
	}

	err = s.store.DeleteUser(s.ctx, users...)
	if err != nil {
		return
	}

	return
}

func (s seeder) LookupModuleByID(ID uint64) (mod *cTypes.Module, err error) {
	if ID == 0 {
		err = fmt.Errorf("invalid ID for module")
		return nil, err
	}
	mod, err = s.store.LookupComposeModuleByID(s.ctx, ID)
	if err != nil {
		return
	}
	return
}

// CreateRecord creates given no of record into DB
func (s seeder) CreateRecord(moduleID uint64, params Params) (IDs []uint64, err error) {
	var records []*cTypes.Record
	var labels []*lTypes.Label

	mod, err := s.LookupModuleByID(moduleID)
	if mod == nil {
		return IDs, fmt.Errorf("module not found")
	}

	if err != nil {
		return nil, err
	}
	for i := 0; i < params.getLimit(); i++ {
		rec := &cTypes.Record{
			ID:          id.Next(),
			NamespaceID: mod.NamespaceID,
			ModuleID:    mod.ID,
			CreatedAt:   time.Now(),
		}

		for i, f := range mod.Fields {
			err := s.setRecordValues(rec, uint(i), f)
			if err != nil {
				return nil, err
			}
		}

		records = append(records, rec)
		labels = append(labels, s.CreateLabel(
			rec.ID,
			rec.LabelResourceKind(),
			fakeDataLabel,
		))
	}

	err = s.store.CreateComposeRecord(s.ctx, mod, records...)
	if err != nil {
		return
	}

	err = s.store.UpsertLabel(s.ctx, labels...)
	if err != nil {
		return
	}

	return
}

// setRecordValues will generate record values from third party and set to record
func (s seeder) setRecordValues(rec *cTypes.Record, place uint, field *cTypes.ModuleField) (err error) {
	// fixme verify all the check and delete
	// figure out what kind of field this is
	//  - type
	//  - name
	//  - is it required? use some kinde of fill ratio? .8
	// - is it multi-value field? how many values?
	// try to get some fake data from the lib
	// create some values

	var value string

	// skip the non required fields
	if !field.Required {
		return
	}

	// skip the non required fields
	if len(field.Name) == 0 {
		return fmt.Errorf("invalid field name")
	}

	// check the type to fake the data
	if len(field.Kind) > 0 {
		value, err = s.faker.fakeValue(field.Name, field.Kind, valueOptions{})
		if err != nil {
			return fmt.Errorf("coudn't generate the value")
		}
	} else {
		return fmt.Errorf("unknown kind for field")
	}

	rec.Values = rec.Values.Set(&cTypes.RecordValue{
		RecordID: rec.ID,
		Name:     field.Name,
		Value:    value,
		// Ref:      0, // @todo we need to talk about this (FK)
		Place: place, // in case of multi-value field this is ++
	})

	return
}

// DeleteAllRecord clear all the fake user from DB
func (s seeder) DeleteAllRecord(mod *cTypes.Module) (err error) {
	filter := cTypes.RecordFilter{
		Labels: map[string]string{
			fakeDataLabel: fakeDataLabel,
		},
	}
	records, _, err := s.store.SearchComposeRecords(s.ctx, mod, filter)
	if err != nil {
		return
	}

	err = s.store.DeleteComposeRecord(s.ctx, mod, records...)
	if err != nil {
		return
	}

	return
}

// DeleteAll will delete all the fake data from the DB
func (s seeder) DeleteAll(mod *cTypes.Module) (err error) {
	err = s.DeleteAllUser()
	if err != nil {
		return
	}

	err = s.DeleteAllRecord(mod)
	if err != nil {
		return err
	}

	return
}
