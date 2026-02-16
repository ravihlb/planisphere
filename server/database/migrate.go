package database

import (
	"database/sql"
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	args := os.Args[1:]

	if len(args) <= 0 {
		fmt.Println("Specify up/down as first arg, then the number of steps as second.")
		fmt.Println("Example: 'up 2'")
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
		up(count)
	}

	if mode == "down" {
		down(count)
	}
}

func up(count int) {
	db, err := sql.Open("sqlite3", "database/planisphere.db")
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

func down(count int) {
	db, err := sql.Open("sqlite3", "database/planisphere.db")
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
