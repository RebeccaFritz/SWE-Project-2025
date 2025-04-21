package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func createLobbyCode(LobbyCode string, wsConnection *websocket.Conn) {
	_, exists := LOBBY[LobbyCode] // Go idiom for checking for existence of a key in a map

	if !exists {
		client := CLIENTS[wsConnection]
		LOBBY[LobbyCode] = client
		return
	}

	badMsg := msgStruct{
		MsgType: "validate lobby code",
		Message: "This lobby code has already been used",
	}
	handleWrite(1, badMsg, wsConnection)
}

// this function handles incoming messages of the type "Lobby Code"
// it (1) checks to see if the provided lobby code is correct,
// (2a) if correct it places both the provided client and the client with the matching code in a new room
// (2b) if wrong it send the client back an error message
func matchLobbyCode(LobbyCode string, wsConnection *websocket.Conn) {
	otherClient, otherClientExists := LOBBY[LobbyCode]
	thisClient, thisClientExists := CLIENTS[wsConnection]

	if(!thisClientExists){
		log.Println("Error: a client tried to join a room but it is not in the client map!")
		return
	}

	if otherClientExists && otherClient.connection != wsConnection {
		room := NewRoom(otherClient, thisClient)
		log.Println("New room: ", room)

		go runGameLoop(false, room)

		goodMsg := msgStruct{
			MsgType: "validate lobby code",
			Message: "Your lobby code has sucessfully matched",
		}
		handleWrite(1, goodMsg, wsConnection)
		handleWrite(1, goodMsg, otherClient.connection) // write confirmation to opponent

		delete(LOBBY, LobbyCode) // remove the used lobby code from the LOBBY map
	} else if otherClientExists && otherClient.connection == wsConnection {
		badMsg := msgStruct{
			MsgType: "validate lobby code",
			Message: "You cannot connect to your own lobby",
		}
		handleWrite(1, badMsg, wsConnection)
	} else {
		badMsg := msgStruct{
			MsgType: "validate lobby code",
			Message: "The provided lobby code does not match any of the existing codes",
		}
		handleWrite(1, badMsg, wsConnection)
	}
}

func NewRoom(client1 *Client, client2 *Client)*Room{
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
	room.inputQueue = []string{}
	// add room to the map
	ROOMS[roomID] = &room

	return &room
}
