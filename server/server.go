package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	db, err := connect_db("/Users/lad/Developer/SWE-Project-2025/data/db.sqlite3")

	if err != nil {
		fmt.Print(err)
		return
	}

	add_user("Amoniker", db)
	add_user("kim", db) // error handling?
	add_user("harry", db)

	result, _ := db.Exec("SELECT username, wins FROM leaderboard ORDER BY wins DESC; ")

	fmt.Print(result)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// add_user adds the given user to the leaderboard table in the database with zero wins.
func add_user(username string, db *sql.DB) error {
	add_user := "INSERT INTO leaderboard (username, wins) VALUES ( ?, 0)"
	_, err := db.Exec(add_user, username)

	return err
}

func connect_db(path string) (*sql.DB, error) {
	// How to update this?
	db, err := sql.Open("sqlite", path)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the SQLite database successfully.")

	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)

	if err != nil {
		return nil, err
	}

	fmt.Println(sqliteVersion)

	return db, err
}

// handle incoming requests and write a response to client
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client!")
}

// the Gameroom struct contains all the information for one game
type Gameroom struct {
	player1token int        // identifier token for the first client in the room
	player2token int        // identifier token for the second client in the room
	targets      [10]Target // struct containing information for each first target
}

// the Target struct
type Target struct {
	twosComp   int    // two's complement number
	baseTen    int    // base 10 number
	hasBoost   bool   // does this target have a boost
	isOnScreen bool   // is this target on screen
	position   [2]int // position as the Target would appear on player 1's screen
}

// the client struct
type Client struct {
	score    int
	health   int    // current health
	position [2]int // position as the token would appear on THIS player's screen
}
