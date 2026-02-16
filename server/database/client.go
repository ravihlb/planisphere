package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func OpenConnection(dsnURI string) *sql.DB {
	db, err := sql.Open("sqlite", dsnURI+"?_foreign_keys=on")

	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	return db
}
