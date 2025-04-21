package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)


// variable used to upgrade HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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
		connection: websocket,
		playerNum:  0, // set to player 0 (will be updated if client joins another room)
	}

	defer closeClient(websocket, client)

	CLIENTS[client.connection] = &client // add client to CLIENTS map

	handleWrite(1, leaderboard, websocket) // write the leaderboard data (1 is the msgType constant for text)
	handleMessaging(websocket)
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
		curRoom, exists := ROOMS[CLIENTS[websocket].roomID]
		if !exists {
			log.Println("Error: Client sent game input but is not in a room")
			break
		}

		// Go automatically dereferences pointers to structures when you access their fields (I was confused by this)
		curRoom.inputQueue = append(curRoom.inputQueue, incomingMsg.Input)
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

func closeClient(websocket *websocket.Conn, client Client) {
	if client.roomID != "" {
		curRoom := ROOMS[client.roomID]
		// remove client from Room by setting it to an uninitialized Client struct
		curRoom.clients[client.playerNum] = &Client{}
	}
	websocket.Close()
}
