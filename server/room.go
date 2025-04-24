// room.go handles
package main

import (
	"github.com/google/uuid"
)

// maps room IDs to room structures
var ROOMS = make(map[string]*Room)

// Room struct contains all the information for one client pair session
type Room struct {
	gamestate  Gamestate
	inputQueue []InputQueueEntry
	clients    [2]*Client // clients in the room
}

// InputQueueEntry defines the elements of an inputQueue
type InputQueueEntry struct {
	input     string
	player    int
}

// NewRoom initializes a new room with two clients, sets default values,
// and adds the room to the global ROOMS map.
func NewRoom(client1 *Client, client2 *Client) *Room {
	roomID := uuid.NewString() // generate unique string to id the room
	room := Room{}

	// place the clients in the room
	room.clients[0] = client1
	room.clients[1] = client2
	// assign player1 and player2
	room.clients[0].playerNum = 1
	room.clients[1].playerNum = 2
	// set default room values
	room.clients[0].roomID = roomID
	room.clients[1].roomID = roomID
	room.gamestate = initGameState()
	room.inputQueue = []InputQueueEntry{}
	// add room to the map
	ROOMS[roomID] = &room

	return &room
}
