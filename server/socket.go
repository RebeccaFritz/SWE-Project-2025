package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// define a global array for the all the lobby codes
var LOBBY = make(map[string]Client)

// define a global array for all the clients (identified by ther websocket connections)
var CLIENTS = make(map[*websocket.Conn]Client)

// the client struct
type Client struct {
	score      int
	health     int             // current health
	position1  [2]int          // position as the token would appear on player 1's screen
	position2  [2]int          // position as the token would appear on player 2's screen
	playerNum  int             // 1 or 2 (default 0)
	roomID     string          // (default "")
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
		roomID:     uuid.NewString(), // generate unique string to id the room (will be updated if client joins another room)
		playerNum:  0,                // set to player 0 (will be updated if client joins another room)
	}

	defer closeClient(websocket, client)

	CLIENTS[client.connection] = client // add client to CLIENTS map

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
	// fmt.Printf("Received: %s\\n", message)

	// decode JSON data with Unmarshal function and store it in a temporary structure
	var newMsgStruct msgStruct
	err = json.Unmarshal(message, &newMsgStruct)
	if err != nil {
		log.Println("Error:", err)
	}

	var curRoom = ROOMS[newMsgStruct.RoomId]

	switch newMsgStruct.MsgType {
	case "client":
		var client = curRoom.clients[newMsgStruct.PlayerNum]
		if newMsgStruct.PlayerNum == 0 { // update client position
			client.position1 = newMsgStruct.Position
			client.position2 = reflect(newMsgStruct.Position)
		} else {
			client.position1 = reflect(newMsgStruct.Position)
			client.position2 = newMsgStruct.Position
		}
	case "target":
		var target = curRoom.targets[newMsgStruct.TargetIdx]
		if newMsgStruct.PlayerNum == 0 { // update target position
			target.position0 = reflect(newMsgStruct.Position)
			target.position1 = newMsgStruct.Position
		} else {
			target.position0 = newMsgStruct.Position
			target.position1 = reflect(newMsgStruct.Position)
		}
	case "create lobby code":
		fmt.Printf("Received: %s\\n", message)
		newLobbyCode := newMsgStruct.lobbyCode
		client := CLIENTS[websocket]
		LOBBY[newLobbyCode] = client
	case "lobby code":
		fmt.Printf("Received: %s\\n", message)
		handleLobbyMessage(newMsgStruct, websocket)
	case "test":
		log.Println("msg: ", newMsgStruct.Message)
	default:
		log.Printf("Error: unknown message type '%s'", newMsgStruct.MsgType)
	}

	return msgType, newMsgStruct, nil
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
	RoomId      string     // the id of the room to which the client who sent the message belongs
	PlayerNum   int        // index of the client in their room (1 or 2)
	TargetIdx   int        // index of the target in the client's room (0 to 9)
	MsgType     string     // the type of msg: "client", "target"
	Position    [2]int     // a target or client position
	Message     string     // other messages
	CurTick     time.Time  // integer messages
	Leaderboard []LB_Entry // array of leaderboard entries
	lobbyCode   string     // for lobby code creation or connection
}

// the reflect function flips the given (x, y) coordinates about the middle of the screen
// so that that object will display correctly on the other player's screen
func reflect(position [2]int) [2]int {
	var reflectedPos = [2]int{-position[0], -position[1]} // flip about the origin
	return reflectedPos
}

func closeClient(websocket *websocket.Conn, client Client) {
	curRoom := ROOMS[client.roomID]
	// remove client from Room by setting it to an uninitialized Client struct
	curRoom.clients[client.playerNum] = Client{}
	websocket.Close()
}

// this function handles incoming messages of the type "Lobby Code"
// it (1) checks to see if the provided lobby code is correct,
// (2a) if correct it places both the provided client and the client with the matching code in a new room
// (2b) if wrong it send the client back an error message
func handleLobbyMessage(message msgStruct, websocket *websocket.Conn) {
	found := false
	for key, value := range LOBBY {
		if message.lobbyCode == key {
			roomID := uuid.NewString() // generate unique string to id the room
			curRoom := ROOMS[roomID]
			// place the clients in the room
			curRoom.clients[0] = value
			curRoom.clients[1] = CLIENTS[websocket]
			// assign player1 and player2
			curRoom.clients[0].playerNum = 1
			curRoom.clients[1].playerNum = 2
			// set default room values
			curRoom.clients[0].score = 0
			curRoom.clients[1].score = 0
			curRoom.clients[0].health = 5
			curRoom.clients[1].health = 5
			curRoom.clients[0].roomID = roomID
			curRoom.clients[1].roomID = roomID
		}
	}

	// if code is not found send out error
	if !found {
		badCodeMessage := msgStruct{
			MsgType: "bad lobby code",
			Message: "The provided lobby code does not match any of the existing codes",
		}
		handleWrite(1, badCodeMessage, websocket)
	}

}
