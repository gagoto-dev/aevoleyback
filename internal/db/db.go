package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error

	db, err = sql.Open("sqlite3", "./app.db") // Open a connection to the SQLite database file named app.db
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS user (
            id UUID PRIMARY KEY,
            name TEXT NOT NULL,
            password TEXT NOT NULL,
            email TEXT NOT NULL,
            createdAt timestamp DEFAULT (now())
        );

        CREATE TABLE IF NOT EXISTS blog_entry (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            authorId UUID NOT NULL,
            title TEXT NOT NULL,

            FOREIGN KEY(authorId) REFERENCES users(id)
        ); `)

	if err != nil {
		return nil, fmt.Errorf("Error creating tables: %s\n", err)
	}

	return db, nil
}
