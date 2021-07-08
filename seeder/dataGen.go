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
	genOption struct {
		totalRecord int
	}
	userGenOption struct {
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

const (
	fakeDataLabel       = "generated"
	fakeUserLabelName   = "generatedUser"
	fakeRecordLabelName = "generatedRecord"
)

func DataGen(ctx context.Context, store store.Storer, faker fakerService) *dataGen {
	return &dataGen{ctx, store, faker}
}

// makeMeFakeDataLabel return the label for generate data
func (gen dataGen) makeMeFakeDataLabel(resourceID uint64, kind, name string) *lTypes.Label {
	return &lTypes.Label{
		Kind:       kind,
		ResourceID: resourceID,
		Name:       name,
		Value:      fakeDataLabel,
	}
}

// makeMeSomeFakeUserPlease creates given no of users into DB
func (gen dataGen) makeMeSomeFakeUserPlease(opt genOption) (IDs []uint64, err error) {
	var userIDs []uint64
	var users []*sTypes.User
	var labels []*lTypes.Label

	for i := 0; i < opt.totalRecord; i++ {
		var user sTypes.User
		user.ID = id.Next()
		user.Email, _ = gen.faker.fakeValueByName("Email")
		user.Name, _ = gen.faker.fakeValueByName("Name")
		user.Handle = gen.faker.fakeUserHandle(user.Name)
		user.Kind = sTypes.BotUser

		userIDs = append(userIDs, user.ID)
		users = append(users, &user)
		labels = append(labels, gen.makeMeFakeDataLabel(
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

// clearFakeUser clear all the fake user from DB
func (gen dataGen) clearFakeUsers() (err error) {
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

// makeMeSomeFakeRecordPlease creates given no of record into DB
func (gen dataGen) makeMeSomeFakeRecordPlease(mod *cTypes.Module) (*cTypes.Record, error) {
	rec := &cTypes.Record{
		ID:          id.Next(),
		NamespaceID: mod.NamespaceID,
		ModuleID:    mod.ID,
		CreatedAt:   time.Now(),
	}

	for i, f := range mod.Fields {
		fmt.Printf("At mod.Fields for loop: %+v", f)
		err := gen.doTheFakeDataMagic(rec, uint(i), f)
		if err != nil {
			return nil, err
		}
	}

	err := gen.store.CreateComposeRecord(gen.ctx, mod, rec)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

// clearFakeUser clear all the fake user from DB
func (gen dataGen) clearFakeRecords() (err error) {
	// fixMe module ID
	filter := cTypes.RecordFilter{
		Labels: map[string]string{
			fakeRecordLabelName: fakeDataLabel,
		},
	}
	records, _, err := gen.store.SearchComposeRecords(gen.ctx, nil, filter)
	if err != nil {
		return
	}

	err = gen.store.DeleteComposeRecord(gen.ctx, nil, records...)
	if err != nil {
		return
	}

	return
}

// fixme:
// better way to return error
// user of method, maybe access third-party pkg from it
func (gen dataGen) doTheFakeDataMagic(rec *cTypes.Record, place uint, field *cTypes.ModuleField) (err error) {
	fmt.Println("\n At doTheFakeDataMagic: ", place, field.Kind)
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

	field.Label = fakeDataLabel

	rec.Values = rec.Values.Set(&cTypes.RecordValue{
		RecordID: rec.ID,
		Name:     field.Name,
		Value:    value,
		// Ref:      0, // @todo we need to talk about this (FK)
		Place: place, // in case of multi-value field this is ++
	})

	fmt.Printf("rec.Values: %+v", rec.Values.String())

	return
}

// clearAllFakeData will delete all the fake data from the DB
func (gen dataGen) clearAllFakeData() (err error) {
	err = gen.clearFakeUsers()
	if err != nil {
		return
	}

	err = gen.clearFakeRecords()
	if err != nil {
		return err
	}

	return
}
