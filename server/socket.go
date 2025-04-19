package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var INPUT_QUEUE = []string{}

// define a global array for the all the lobby codes
var LOBBY = make(map[string]Client)

// define a global array for all the clients (identified by ther websocket connections)
var CLIENTS = make(map[*websocket.Conn]Client)

// define an empty Client struct for refrence purposes
var zeroValClient Client

// the client struct
type Client struct {
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
		playerNum:  0, // set to player 0 (will be updated if client joins another room)
	}

	defer closeClient(websocket, client)

	go runGameLoop(false, client, client)

	CLIENTS[client.connection] = client // add client to CLIENTS map

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

	// var curRoom = ROOMS[incomingMsg.RoomId]

	switch incomingMsg.MsgType {
		case "create lobby code":
			createLobbyCode(incomingMsg.LobbyCode, websocket)
		case "lobby code":
			matchLobbyCode(incomingMsg.LobbyCode, websocket)
		case "test":
			log.Println("msg: ", incomingMsg.Message)
		case "input":
			INPUT_QUEUE = append(INPUT_QUEUE, incomingMsg.Input)
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

// this struct temporarily stores incoming message data before it is validated (if it starts with an uppercase letter it can be exported by Marshal())
type msgStruct struct {
	RoomId      string     // the id of the room to which the client who sent the message belongs
	PlayerNum   int        // index of the client in their room (1 or 2)
	TargetIdx   int        // index of the target in the client's room (0 to 9)
	MsgType     string     // the type of msg: "client", "target"
	Message     string     // other messages
	CurTick     time.Time  // integer messages
	Leaderboard []LB_Entry // array of leaderboard entries
	Gamestate 	Gamestate
	Input string
	LobbyCode   string     // for lobby code creation or connection
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

func createLobbyCode(LobbyCode string, wsConnection *websocket.Conn) {
	value := LOBBY[LobbyCode]

	if value == zeroValClient {
		// if the lobby code has not been used
		client := CLIENTS[wsConnection]
		LOBBY[LobbyCode] = client
	} else {
		badMsg := msgStruct{
			MsgType: "validate lobby code",
			Message: "This lobby code has already been used",
		}
		handleWrite(1, badMsg, wsConnection)
	}
}

// this function handles incoming messages of the type "Lobby Code"
// it (1) checks to see if the provided lobby code is correct,
// (2a) if correct it places both the provided client and the client with the matching code in a new room
// (2b) if wrong it send the client back an error message
func matchLobbyCode(LobbyCode string, wsConnection *websocket.Conn) {

	value := LOBBY[LobbyCode]

	if value != zeroValClient && value.connection != wsConnection {
		roomID := uuid.NewString() // generate unique string to id the room
		curRoom := ROOMS[roomID]
		// place the clients in the room
		curRoom.clients[0] = value
		curRoom.clients[1] = CLIENTS[wsConnection]
		// assign player1 and player2
		curRoom.clients[0].playerNum = 1
		curRoom.clients[1].playerNum = 2
		// set default room values
		curRoom.clients[0].roomID = roomID
		curRoom.clients[1].roomID = roomID

		goodMsg := msgStruct{
			MsgType: "validate lobby code",
			Message: "Your lobby code has sucessfully matched",
		}
		handleWrite(1, goodMsg, wsConnection)
		handleWrite(1, goodMsg, value.connection) // write confirmation to opponent

		delete(LOBBY, LobbyCode) // remove the used lobby code from the LOBBY map
	} else if value != zeroValClient && value.connection == wsConnection {
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
