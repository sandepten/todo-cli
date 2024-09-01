package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func New() (*Database, error) {
	// if the sqlite database does not exist, it will be created
	if _, err := os.Stat("./todo.db"); os.IsNotExist(err) {
		file, err := os.Create("./todo.db")
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return nil, err
	}

	// create the tables
	if err := (&Database{db: db}).createTables(); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

// CreateTables creates the tables for the todo list
func (d *Database) createTables() error {
	// user table
	_, err := d.db.Exec(`CREATE TABLE IF NOT EXISTS user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE
    )`)
	if err != nil {
		log.Println("Error creating user table:", err)
		return err
	}

	return nil
}
