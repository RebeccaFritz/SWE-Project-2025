package main

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

// 'add_user' adds the given user to the leaderboard table in the given db.
func add_user(username string, db *sql.DB) error {
	add_user_sql := "INSERT INTO leaderboard (username, wins) VALUES ( ?, 0)"
	_, err := db.Exec(add_user_sql, username)

	return err
}

// 'increment_wins' increments the wins of the given user in the leaderboard table in the db.
func increment_wins(username string, db *sql.DB) error {
	increment_wins_sql := "UPDATE leaderboard SET wins = wins + 1 WHERE username = ?;"
	_, err := db.Exec(increment_wins_sql, username)

	return err
}

// Allows server code to query the database.
func connect_db(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Connected to the SQLite database successfully.")

	// get the sqlite version and print it.
	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)

	if err != nil {
		fmt.Print()
		return nil, err
	}

	fmt.Println(sqliteVersion)
	return db, err
}
