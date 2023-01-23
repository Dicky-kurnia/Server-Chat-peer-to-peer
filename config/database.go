package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToSQLDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=chat_app password=password sslmode=disable")
	if err != nil {
		return nil, err
	}
	AutoMigrate()
	return db, nil
}

func AutoMigrate() {
	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);`)

	db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		sender_id INTEGER REFERENCES users(id),
		receiver_id INTEGER REFERENCES users(id),
		message TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
}
