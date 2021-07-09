package seeder

import (
	"github.com/brianvoe/gofakeit/v6"
	"strings"
)

type (
	faker struct {
		methods map[string]func() string
	}

	valueOptions struct {
	}
)

var (
	fakerMethods = map[string]func() string{
		"Name":        gofakeit.Name,
		"FirstName":   gofakeit.FirstName,
		"LastName":    gofakeit.LastName,
		"Title":       gofakeit.JobTitle,
		"Phone":       gofakeit.Phone,
		"MobilePhone": gofakeit.Phone,
		"Email":       gofakeit.Email,
	}
)

func Faker() *faker {
	return &faker{fakerMethods}
}

// seed to ensure randomization on initial
func (f faker) seed() {
	gofakeit.Seed(0)
	return
}

// generateValue generate value based on name or given type
func (f faker) fakeValueByName(name string) (val string, ok bool) {
	// Ensure randomization on initial
	f.seed()

	// Generate & return value from mapped methods
	method, ok := f.methods[name]
	if ok {
		return method(), ok
	}
	return
}

// generateValue generate value based on name or given type
func (f faker) fakeValue(name, kind string, opt valueOptions) (val string, err error) {
	// Generate & return value from mapped methods
	val, ok := f.fakeValueByName(name)
	if ok {
		return
	}

	// Ensure randomization on initial
	f.seed()
	valueKind := toLowerCase(kind)

	// Since, we don't have faker method for it,
	// we will generate the value based on kind(type)
	switch true {
	case valueKind == "string":
		val = gofakeit.LoremIpsumWord()
		break
	case valueKind == "int":
		break
	}
	return
}

func (f faker) fakeUserHandle(s string) string {
	return gofakeit.Generate("??????") + "_seeded"
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}
