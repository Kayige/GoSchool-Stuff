package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type connection struct{}

var (
	con  *connection
	once sync.Once
)

// GetInstance return the database instance
func GetInstance() *connection {
	// 1 intance form stuct connection
	once.Do(func() {
		err := godotenv.Load()

		if err != nil {
			log.Fatal(err)
		}

		con = &connection{}
	})
	return con
}

// GetConnection return db connection
func (c *connection) GetConnection() *gorm.DB {

	connArgs := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("GO_ENV") == "development" {
		db.LogMode(true)
	}

	if err != nil {
		log.Fatal(err)
	}

	return db
}
