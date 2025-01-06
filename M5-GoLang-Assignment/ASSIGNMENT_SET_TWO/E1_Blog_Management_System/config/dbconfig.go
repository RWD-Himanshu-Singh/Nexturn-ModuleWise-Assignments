package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitializeDatabase() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite", "./blog.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
		return nil, err
	}
	return DB, nil
}
