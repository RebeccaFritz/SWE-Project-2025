package main

import (
	"database/sql"
	"fmt"
	"log"

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

// get the leaderboard data from the SQL database
func getLeaderboard(db *sql.DB) msgStruct {
	sql := "Select * FROM leaderboard"
	rows, err := db.Query(sql) // execute the select statement and return a single row
	if err != nil {
		log.Println("Error querying table: ", err)
	}
	defer rows.Close()

	// query each row and save the data in the lb struct
	var entries []LB_Entry
	for rows.Next() {
		entry := &LB_Entry{}
		err := rows.Scan(&entry.Username, &entry.Wins)
		if err != nil {
			log.Println("Error querying rows: ", err)
		}
		entries = append(entries, *entry)
	}

	// return a msgStruct with all the leaderboard entries
	return msgStruct{MsgType: "leaderboard", Leaderboard: entries}
}

// a single leaderboard entry
type LB_Entry struct {
	Username string // username of the player
	Wins     int    // number of wins
}

// RUN ONCE
func create_table(db *sql.DB) error {
	create_table_sql := "CREATE TABLE leaderboard (username TEXT PRIMARY KEY,  wins INTEGER);"
	_, err := db.Exec(create_table_sql)

	return err
}
