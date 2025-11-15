package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "blog.db")
	if err != nil {
		panic("Could'nt connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(3)

	createTables()
}

func createTables() {
	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	createBlogs := `
	CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author_id INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(author_id) REFERENCES users(id)
	);
	`

	_, err := DB.Exec(createUsers)
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec(createBlogs)
	if err != nil {
		panic(err)
	}
}
