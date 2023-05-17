package models

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDataBase() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "url-shortener-db",
	}
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", pingErr)
	}
	fmt.Println("We are connected to the database ")
}
