package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SetupModels load Schemma into DB
func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "local.db")

	if err != nil {
		log.Fatal(err)
		panic("Failed connect to Database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Note{})

	return db
}
