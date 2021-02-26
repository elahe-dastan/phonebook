package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //adding dialect for postgres
)

func New() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=sotoon password=postgres sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}

	return db
}
