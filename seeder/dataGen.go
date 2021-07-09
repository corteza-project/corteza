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
	dataGen struct {
		ctx   context.Context
		store storeService
		faker fakerService
	}
	GenOption struct {
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
		SearchComposeRecords(ctx context.Context, _mod *cTypes.Module, f cTypes.RecordFilter) (cTypes.RecordSet, cTypes.RecordFilter, error)
		CreateComposeRecord(ctx context.Context, mod *cTypes.Module, rr ...*cTypes.Record) error
		DeleteComposeRecord(ctx context.Context, m *cTypes.Module, rr ...*cTypes.Record) error
	}
)

var (
	DefaultStore store.Storer
)

const (
	fakeDataLabel       = "generated"
	fakeUserLabelName   = "generatedUser"
	fakeRecordLabelName = "generatedRecord"
)

func DataGen(ctx context.Context, store store.Storer, faker fakerService) *dataGen {
	DefaultStore = store
	return &dataGen{ctx, store, faker}
}

// getLimit return data generation limit; It will return Default(1) if limit is 0
func (gOpt GenOption) getLimit() int {
	if gOpt.Limit == 0 {
		return 1
	}
	return gOpt.Limit
}

// MakeMeFakeDataLabel return the label for generate data
func (gen dataGen) MakeMeFakeDataLabel(resourceID uint64, kind, name string) *lTypes.Label {
	return &lTypes.Label{
		Kind:       kind,
		ResourceID: resourceID,
		Name:       name,
		Value:      fakeDataLabel,
	}
}

// MakeMeSomeFakeUserPlease creates given no of users into DB
func (gen dataGen) MakeMeSomeFakeUserPlease(opt GenOption) (IDs []uint64, err error) {
	var users []*sTypes.User
	var labels []*lTypes.Label

	for i := 0; i < opt.getLimit(); i++ {
		var user sTypes.User
		user.ID = id.Next()
		user.Email, _ = gen.faker.fakeValueByName("Email")
		user.Name, _ = gen.faker.fakeValueByName("Name")
		user.Handle = gen.faker.fakeUserHandle(user.Name)
		user.Kind = sTypes.BotUser
		user.CreatedAt = time.Now()

		IDs = append(IDs, user.ID)
		users = append(users, &user)
		labels = append(labels, gen.MakeMeFakeDataLabel(
			user.ID,
			user.LabelResourceKind(),
			fakeUserLabelName,
		))
	}

	err = gen.store.CreateUser(gen.ctx, users...)
	if err != nil {
		return
	}

	err = gen.store.UpsertLabel(gen.ctx, labels...)
	if err != nil {
		return
	}

	return
}

// ClearFakeUsers clear all the fake user from DB
func (gen dataGen) ClearFakeUsers() (err error) {
	filter := sTypes.UserFilter{
		Labels: map[string]string{
			fakeUserLabelName: fakeDataLabel,
		},
	}
	users, _, err := gen.store.SearchUsers(gen.ctx, filter)
	if err != nil {
		return
	}

	err = gen.store.DeleteUser(gen.ctx, users...)
	if err != nil {
		return
	}

	return
}

// MakeMeSomeFakeRecordPlease creates given no of record into DB
func (gen dataGen) MakeMeSomeFakeRecordPlease(mod *cTypes.Module, opt GenOption) (IDs []uint64, err error) {
	var records []*cTypes.Record
	var labels []*lTypes.Label

	for i := 0; i < opt.getLimit(); i++ {
		rec := &cTypes.Record{
			ID:          id.Next(),
			NamespaceID: mod.NamespaceID,
			ModuleID:    mod.ID,
			CreatedAt:   time.Now(),
		}

		for i, f := range mod.Fields {
			err := gen.doTheFakeDataMagic(rec, uint(i), f)
			if err != nil {
				return nil, err
			}
		}

		records = append(records, rec)
		labels = append(labels, gen.MakeMeFakeDataLabel(
			rec.ID,
			rec.LabelResourceKind(),
			fakeRecordLabelName,
		))
	}

	err = gen.store.CreateComposeRecord(gen.ctx, mod, records...)
	if err != nil {
		return
	}

	err = gen.store.UpsertLabel(gen.ctx, labels...)
	if err != nil {
		return
	}

	return
}

// ClearFakeRecords clear all the fake user from DB
func (gen dataGen) ClearFakeRecords(mod *cTypes.Module) (err error) {
	filter := cTypes.RecordFilter{
		Labels: map[string]string{
			fakeRecordLabelName: fakeDataLabel,
		},
	}
	records, _, err := gen.store.SearchComposeRecords(gen.ctx, mod, filter)
	if err != nil {
		return
	}

	err = gen.store.DeleteComposeRecord(gen.ctx, mod, records...)
	if err != nil {
		return
	}

	return
}

// fixme:
// better way to return error
// user of method, maybe access third-party pkg from it
func (gen dataGen) doTheFakeDataMagic(rec *cTypes.Record, place uint, field *cTypes.ModuleField) (err error) {
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
		value, err = gen.faker.fakeValue(field.Name, field.Kind, valueOptions{})
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

// ClearAllFakeData will delete all the fake data from the DB
func (gen dataGen) ClearAllFakeData(mod *cTypes.Module) (err error) {
	err = gen.ClearFakeUsers()
	if err != nil {
		return
	}

	err = gen.ClearFakeRecords(mod)
	if err != nil {
		return err
	}

	return
}
