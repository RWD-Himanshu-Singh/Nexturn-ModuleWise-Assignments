package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create Products table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		stock INTEGER,
		category_id INTEGER
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return DB
}
