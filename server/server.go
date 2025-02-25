package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
)

// use go run *.go, otherwise it wont run all the files

func main() {
	db, err := connect_db("/Users/lad/Developer/SWE-Project-2025/data/db.sqlite3")

	if err != nil {
		fmt.Print(err)
		return
	}

	// test calls
	add_user("Amoniker", db)
	add_user("kim", db)
	add_user("harry", db)
	increment_wins("Amoniker", db)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
