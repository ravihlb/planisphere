package main

import (
	"log"
	"planisphere/server/database"

	_ "modernc.org/sqlite"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db := database.OpenConnection("./database/planisphere.db")
	defer db.Close()
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	// defer cancel()

	log.Println("Querying...")

	rows, err := db.Query("select * from sqlite_schema;")
	if err != nil {
		log.Fatal(err)
	}

}

// TODO
//
// Set up connection to SQLite DB
// Build API Spec
// Serve API
//  - Set up basic HTTP server
//  - Set up auth
//  - Set up JSON responding following the API spec
