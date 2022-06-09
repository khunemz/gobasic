package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("mydb.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect database ! \n ", err.Error())
		os.Exit(2)
	}
	log.Printf("Connected to database successfully! \n")
	db.Logger = logger.Default.LogMode(logger.Info).LogMode(logger.Warn).LogMode(logger.Error)
	log.Printf("Database migration goes here \n")

	Database := DbInstance{Db: db}
}
