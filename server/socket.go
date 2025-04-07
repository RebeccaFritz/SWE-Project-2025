package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// the client struct
type Client struct {
	score      int
	health     int    // current health
	position0  [2]int // position as the token would appear on player 0's screen
	position1  [2]int // position as the token would appear on player 1's screen
	playerNum  int    // 0 or 1
	roomID     string
	connection *websocket.Conn // the websocket connection this client is on
}

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

	handleWrite(1, leaderboard, client.connection) // write the leaderboard data (1 is the msgType constant for text)
	handleMessaging(websocket, client)
}

// a function that sends a message to a single client
func handleMessaging(websocket *websocket.Conn, client Client) {
	for tick := range time.Tick(time.Second / 1000) {
		// the read waits until a message is recieved
		msgType, message, err := handleRead(websocket)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		message.CurTick = tick

		handleWrite(msgType, message, client.connection) // echo back message
	}
}

// handleRead reads an incoming JSON message from a client and parses it
func handleRead(websocket *websocket.Conn) (int, msgStruct, error) {
	msgType, message, err := websocket.ReadMessage()
	if err != nil {
		return msgType, msgStruct{}, err
	}
	fmt.Printf("Received: %s\n", message)

	// decode JSON data with Unmarshal function and store it in a temporary structure
	var msgStruct msgStruct
	err = json.Unmarshal(message, &msgStruct)
	if err != nil {
		log.Println("Error:", err)
	}

	var curRoom = ROOMS[msgStruct.roomId]

	switch msgStruct.MsgType {
	case "client":
		var client = curRoom.clients[msgStruct.playerIdx]
		if msgStruct.playerIdx == 0 { // update client position
			client.position0 = msgStruct.Position
			client.position1 = reflect(msgStruct.Position)
		} else {
			client.position0 = reflect(msgStruct.Position)
			client.position1 = msgStruct.Position
		}
	case "target":
		var target = curRoom.targets[msgStruct.targetIdx]
		if msgStruct.playerIdx == 0 { // update target position
			target.position0 = reflect(msgStruct.Position)
			target.position1 = msgStruct.Position
		} else {
			target.position0 = msgStruct.Position
			target.position1 = reflect(msgStruct.Position)
		}
	default:
		log.Println("Error: unknown message type")
	}

	return msgType, msgStruct, nil
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

// this struct temporarily stores incoming message data before it is validated (if it starts with an uppercase letter it can be exported by Marshal())
type msgStruct struct {
	roomId      string     // the id of the room to which the client who sent the message belongs
	playerIdx   int        // index of the client in their room (0 or 1)
	targetIdx   int        // index of the target in the client's room (0 to 9)
	MsgType     string     // the type of msg: "client", "target"
	Position    [2]int     // a target or client position
	Message     string     // other messages
	CurTick     time.Time  // integer messages
	Leaderboard []LB_Entry // array of leaderboard entries
}

// the reflect function flips the given (x, y) coordinates about the middle of the screen
// so that that object will display correctly on the other player's screen
func reflect(position [2]int) [2]int {
	var reflectedPos = [2]int{-position[0], -position[1]} // flip about the origin
	return reflectedPos
}

func closeClient(websocket *websocket.Conn, client Client) {
	if client.roomID != "" {
		curRoom := ROOMS[client.roomID]
		// remove client from Room by setting it to an uninitialized Client struct
		curRoom.clients[client.playerNum] = Client{}
	}
	websocket.Close()
}
