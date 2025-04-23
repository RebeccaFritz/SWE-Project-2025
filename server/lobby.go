// lobby.go handles lobby code creation, validation, and matchmaking between clients.
package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// maps lobby codes to clients
var LOBBY = make(map[string]*Client)

// createLobbyCode registers a new lobby code if it doesn't already exist;
// otherwise, it notifies the client that the code is taken.
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

// matchLobbyCode handles "Lobby Code" messages by checking if the code matches an existing lobby.
// If valid and not self-matching, it creates a new room and starts the game.
// If invalid or self-matching, it sends an appropriate error message to the client.
func matchLobbyCode(LobbyCode string, wsConnection *websocket.Conn) {
	otherClient, otherClientExists := LOBBY[LobbyCode]
	thisClient, thisClientExists := CLIENTS[wsConnection]

	if !thisClientExists {
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
