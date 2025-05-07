package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// variable used to upgrade HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type socket struct {
	websocket 	*websocket.Conn
	sync.Mutex
}


// function to handle WebSocket connections
func wsHandler(writer http.ResponseWriter, request *http.Request) {
	// Upgrade HTTP connection to WebSocket connection
	websocket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("Error upgrading:", err)
		return
	}

	// create a new client structure with this websocket connection
	client := Client{
		connection: socket{websocket, sync.Mutex{}},
		playerNum:  0, // set to player 0 (will be updated if client joins another room)
	}

	defer closeClient(websocket, &client)

	CLIENTS[client.connection.websocket] = &client // add client to CLIENTS map

	handleWrite(1, LEADERBOARD, &client.connection) // write the leaderboard data (1 is the msgType constant for text)
	handleMessaging(&client.connection)
}

func closeClient(websocket *websocket.Conn, client *Client) {
	if client.roomID != "" {
		curRoom := ROOMS[client.roomID]
		// remove client from Room by setting it to an uninitialized Client struct
		curRoom.clients[client.playerNum-1] = &Client{}
	}
	websocket.Close()
}
