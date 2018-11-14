package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// D is our db connection object
var DB *gorm.DB
var err error

// Connect connects to mysql database and runs migrations
func Connect(un string, pw string, nDb string) {
	// Connect to mysql
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", un, pw, nDb))
	if err != nil {
		log.Fatal(err)
	}
}

// Close closes database connection
func Close() {
	DB.Close()
}
