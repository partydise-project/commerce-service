package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
In order to have a single connection to the database,
the Singleton design pattern is used to initialize
the *gorm.DB object only once.
*/
var DB *gorm.DB
var once sync.Once

func InicializeDB(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	/*With GORM, it is not necessary to use a defer statement
	to close the connection, as the library handles this automatically.*/

	/*The synchronization with sync.Once ensures that the
	initialization of the *DB object only happens once. */
	once.Do(func() {
		DB = db
		errorMigraciones := db.AutoMigrate(&CategoriaEvento{}, &PlanEvento{})
		if errorMigraciones != nil {
			err = fmt.Errorf("error occurred while performing database migrations: %w", errorMigraciones)
			return
		}
	})

	return nil
}
