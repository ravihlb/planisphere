package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// flagging, clearly
	migrationsDirectory := *flag.String("migrations directory", "", "pass the path to the directory where your migration files are located")
	dbFilepath := *flag.String("database file path", "", "database file path")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	args := os.Args[2:]
	flag.Parse()

	if len(args) <= 0 || len(migrationsDirectory) <= 0 || len(dbFilepath) <= 0 {
		fmt.Println("Pass migrations and database flags with ")
		fmt.Println("Example: './database/test.db up 2'")
		os.Exit(1)
	}

	mode := args[0]
	count, err := strconv.Atoi(args[1])

	if err != nil {
		log.Fatal(err)
	}

	if mode != "up" && mode != "down" {
		log.Fatalln("Please provide both mode and count arguments")
	}

	if mode == "up" {
		up(dbFilepath, count)
	}

	if mode == "down" {
		down(dbFilepath, count)
	}
}

func up(dbFilepath string, count int) {
	db, err := sql.Open("sqlite3", dbFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "sqlite3", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Steps(count); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Printf("Migrated up by %d steps\n", count)
}

func down(dbFilepath string, count int) {
	db, err := sql.Open("sqlite3", dbFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "sqlite3", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Steps(-count); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Printf("Migrated down by %d steps\n", count)
}
