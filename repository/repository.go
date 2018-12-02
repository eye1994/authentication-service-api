package repository

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgress addapter used by gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // dev only
)

// DB holds the connection to the database
var DB *gorm.DB
var err error

func init() {
	switch os.Getenv("ENV") {
	case "DEV":
		DB, err = gorm.Open("sqlite3", "dev.db")
	default:
		DB, err = gorm.Open("sqlite3", "test.db")
	}

	if err != nil {
		fmt.Printf("%v\n", err)
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Application{})
	DB.AutoMigrate(&Administrator{})
	DB.AutoMigrate(&User{})
}
