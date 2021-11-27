package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

    // Connection string yaratılır
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	
	
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	//Database migration
	db.Debug().AutoMigrate(&Currency{})
	// db.Debug().AutoMigrate(&CurrencyHistory{})
	// db.Debug().AutoMigrate(&History{})
	// db.Debug().AutoMigrate(&Price{})
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}