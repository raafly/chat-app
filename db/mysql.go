package db

import (
	"database/sql"
	"log"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:saturna@tcp(127.0.0.1)/realtime")
	if err != nil {
		log.Printf("error start db %v", err)
	}

	db.SetConnMaxIdleTime(5 *time.Minute)
	db.SetConnMaxLifetime(60 *time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return db
}