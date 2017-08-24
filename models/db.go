package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:@/api_products?charset=utf8")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
