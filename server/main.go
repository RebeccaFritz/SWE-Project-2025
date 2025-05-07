package main

import (
	"fmt"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	db, err := connect_db("../data/db.sqlite")
	DB = db

	if err != nil {
		fmt.Print(err)
		return
	}
	defer DB.Close()

	create_table(db) // uncomment this if you have never built the app before

	// get leaderboard data from SQL database
	LEADERBOARD = getLeaderboard(DB)

	http.Handle("/", http.FileServer(http.Dir("./dist")))
	// os.Create("../server/server-messages.txt") // create a file to recieve incoming messages to the server
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket server started on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
			fmt.Println("Server error:", err)
	}
}
