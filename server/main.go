package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	db, err := connect_db("../data/db.sqlite")

	if err != nil {
		fmt.Print(err)
		return
	}
	defer db.Close()

	//create_table(db) // uncomment this if you have never built the app before

	// test calls
	// add_user("Amoniker", db)
	// add_user("kim", db)
	// add_user("harry", db)
	// increment_wins("Amoniker", db)

	// get leaderboard data from SQL database
	LEADERBOARD = getLeaderboard(db)

	os.Create("../server/server-messages.txt") // create a file to recieve incoming messages to the server
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket server started on :8080")
	http.ListenAndServe(":8080", nil)
}
