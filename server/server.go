package main

import (
	"fmt"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	db, err := connect_db("../data/db.sqlite")

	if err != nil {
		fmt.Print(err)
		return
	}

	// test calls
	add_user("Amoniker", db)
	add_user("kim", db)
	add_user("harry", db)
	increment_wins("Amoniker", db)

	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket server started on :8080")
	http.ListenAndServe(":8080", nil)
}
