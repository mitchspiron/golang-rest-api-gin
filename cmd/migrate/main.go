package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	// check command line arguments
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: up or down")
	}

	direction := os.Args[1]
	
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Create a new SQLite3 driver instance
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	
	if err != nil {
		log.Fatal(err)
	}

	// Use file source for migrations
	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")
	
	if err != nil {
		log.Fatal(err)
	}

	// Create migrate instance with file source and sqlite3 database
	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)

	if err != nil {
		log.Fatal(err)
	}

	// Apply or rollback migrations based on the direction argument
	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations applied successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations rolled back successfully.")
	default:
		log.Fatal("Unknown direction. Use 'up' or 'down'.")
	}
}