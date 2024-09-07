package database

import (
	"log"
)

func (d *Database) UserExists(name string) (bool, error) {
	var exists bool
	err := d.db.QueryRow("SELECT EXISTS(SELECT 1 FROM user WHERE name = ?)", name).Scan(&exists)
	if err != nil {
		log.Println("Error checking if user exists:", err)
	}
	return exists, err
}
