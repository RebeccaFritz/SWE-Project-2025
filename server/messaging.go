// messaging.go handles reading and writing to clients once a socket has been created
package main

import (
	"encoding/json"
	"log"
	"time"
	"os"
	"fmt"

	"github.com/gorilla/websocket"
)

// this struct temporarily stores incoming message data before it is validated (if it starts with an uppercase letter it can be exported by Marshal())
type msgStruct struct {
	RoomId      string     // the id of the room to which the client who sent the message belongs
	PlayerNum   int        // index of the client in their room (1 or 2)
	TargetIdx   int        // index of the target in the client's room (0 to 9)
	MsgType     string     // the type of msg: "client", "target"
	Message     string     // other messages
	CurTick     time.Time  // integer messages
	Leaderboard []LB_Entry // array of leaderboard entries
	Gamestate   Gamestate
	Input       string
	LobbyCode   string // for lobby code creation or connection
}

// a function that sends a message to a single client
func handleMessaging(wsConnection *websocket.Conn) {

	// according to Prof. Mirabelli, the ideal tick rates are sec/30, sec/60, and sec/120
	for tick := range time.Tick(time.Second / 60) {
		// the read waits until a message is recieved
		msgType, message, err := handleRead(wsConnection)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		message.CurTick = tick

		handleWrite(msgType, message, wsConnection) // echo back message
	}
}

// handleWrite writes a message to a client
func handleWrite(msgType int, msgStruct msgStruct, websocket *websocket.Conn) {
	message, err := json.Marshal(msgStruct)
	if err != nil {
		log.Println("Error Marshaling message:", err)
	}
	err = websocket.WriteMessage(msgType, message)
	if err != nil {
		log.Println("Error writing message:", err)
	}
}

// handleRead reads an incoming JSON message from a client and parses it
func handleRead(websocket *websocket.Conn) (int, msgStruct, error) {
	msgType, message, err := websocket.ReadMessage()
	if err != nil {
		return msgType, msgStruct{}, err
	}

	// write the recieved message to a file
	file, err := os.OpenFile("../server/server-messages.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	clientName := "Client, Local Address: " + websocket.LocalAddr().String() + ", Remote Address: " + websocket.RemoteAddr().String()
	date := "Date: " + time.Now().String()
	received := "Received: " + string(message)
	space := " "
	fullMessage := clientName + "\n" + date + "\n" + received + "\n" + space

	writeToFile(file, fullMessage)

	// decode JSON data with Unmarshal function and store it in a temporary structure
	var incomingMsg msgStruct
	err = json.Unmarshal(message, &incomingMsg)
	if err != nil {
		log.Println("Error:", err)
	}

	switch incomingMsg.MsgType {
	case "create lobby code":
		createLobbyCode(incomingMsg.LobbyCode, websocket)
	case "lobby code":
		matchLobbyCode(incomingMsg.LobbyCode, websocket)
	case "test":
		log.Println("Test msg: ", incomingMsg.Message)
	case "status":
		log.Println("Client Status: ", incomingMsg.Message)
	case "input":
		client, exists := CLIENTS[websocket]
		if !exists {
			log.Println("Error: Recieved game input from a socket, but the socket is not mapped to a client struct")
			break
		}

		room, exists := ROOMS[client.roomID]
		if !exists {
			log.Println("Error: Client sent game input but is not in a room")
			break
		}

		room.inputQueue = append(room.inputQueue, incomingMsg.Input)
	default:
		log.Printf("Error: unknown message type '%s'", incomingMsg.MsgType)
	}

	return msgType, incomingMsg, nil
}

// function to write a message to a file
func writeToFile(file *os.File, message string) {
	_, err := fmt.Fprintln(file, message)
	if err != nil {
		log.Println(err)
	}
}
