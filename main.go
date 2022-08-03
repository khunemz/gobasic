package main

import (
	"fmt"
	"go-hexagonal/handler"
	"go-hexagonal/logs"
	"go-hexagonal/repository"
	"go-hexagonal/service"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func main() {

	// use gorm orm
	// db, err := gorm.Open(sqlite.Open("Database.db"), &gorm.Config{})
	// db.AutoMigrate(&repository.Customer{})
	initConfig()
	initTimeZone()
	db := initDatabase()
	// use sqlx

	// /* migration script (for sqlite only) */
	// schema := `CREATE TABLE customers (customer_id integer not null primary key AUTOINCREMENT, name varchar(100) not null,
	// date_of_birth varchar(30) null, city varchar(30) null, zip_code varchar(30) null, status tinyint not null default 1);`
	// result, migrationErr := db.Exec(schema)
	// if migrationErr != nil {
	// 	panic(migrationErr)
	// }
	// fmt.Println("migration result : ", result)
	// db.MustExec("DELETE FROM customers")
	/* seedint data */
	// db.MustExec("INSERT INTO customers ( name, date_of_birth, city, zip_code, status) VALUES (?, ?, ?, ?, ?)",
	// 	"Chaiwat", "1992-01-13", "Bangkok", "10110", 1)
	// db.MustExec("INSERT INTO customers ( name, date_of_birth, city, zip_code, status) VALUES (?, ?, ?, ?, ?)",
	// 	"Chutipong", "1990-05-15", "Bangkok", "12150", 1)
	// db.MustExec("INSERT INTO customers ( name, date_of_birth, city, zip_code, status) VALUES (?, ?, ?, ?, ?)",
	// 	"Uychai", "1993-02-03", "Bangkok", "11001", 1)
	// db.MustExec("INSERT INTO customers ( name, date_of_birth, city, zip_code, status) VALUES (?, ?, ?, ?, ?)",
	// 	"Mark", "1994-01-12", "Bangkok", "10150", 1)
	// db.MustExec("INSERT INTO customers ( name, date_of_birth, city, zip_code, status) VALUES (?, ?, ?, ?, ?)",
	// 	"Jenet", "1992-10-11", "Bangkok", "10120", 1)
	// db.MustExec("INSERT INTO customers ( name, date_of_birth, city, zip_code, status) VALUES (?, ?, ?, ?, ?)",
	// 	"Robert", "1989-05-01", "Bangkok", "10130", 1)
	// db.MustExec("INSERT INTO customers ( name, date_of_birth, city, zip_code, status) VALUES (?, ?, ?, ?, ?)",
	// 	"Lotus", "1998-02-28", "Bangkok", "10140", 1)

	/** layers */
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryMock
	custService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(custService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	port := viper.GetInt("app.port")
	logs.Info("Start app in port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// check config whether or not existing if not it will overide
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		panic("Cannot read config file")
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%s", viper.GetString("db.host"))
	db, err := sqlx.Connect(viper.GetString("db.driver"), dsn)

	if err != nil {
		logs.Error(fmt.Sprintf("Error connot connect to database ... %v", err))
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	return db
}
