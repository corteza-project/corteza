package seeder

import (
	"database/sql"
	"go.uber.org/zap"
	"log"
	"reflect"
)

type (
	seeder struct {
		logger *zap.Logger
	}

	// fixme: move to seeder
	// temp fix: got to use store layer
	seed struct {
		db *sql.DB
	}
)

func Seeder(logger *zap.Logger) *seeder {
	// todo: get the config, do logEnable check;
	return &seeder{
		logger: logger.Named("seeder"),
	}
}

// Fixme log, fmt,.. => logger

// Execute will executes the given seeder method
func Execute(db *sql.DB, seedMethodNames ...string) {
	s := seed{db}

	seedType := reflect.TypeOf(s)

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			Seed(s, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		Seed(s, item)
	}
}

func Seed(s seed, seedMethodName string) {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succeed")
}
