package main

import (
	"github.com/gorilla/websocket"
	"time"
)

var leaderboard msgStruct // make leaderboard global

// maps lobby codes to clients
var LOBBY = make(map[string]*Client)

// maps websockets to clients
var CLIENTS = make(map[*websocket.Conn]*Client)

// maps room IDs to room structures
var ROOMS = make(map[string]*Room)

// define an empty Client struct for refrence purposes
var zeroValClient Client

// the client struct
type Client struct {
	playerNum  int             // 1 or 2 (default 0)
	roomID     string          // (default "")
	connection *websocket.Conn // the websocket connection this client is on
}

// the Room struct contains all the information for one game
type Room struct {
	gamestate  Gamestate
	inputQueue []string
	clients    [2]*Client // clients in the room
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
	Gamestate   Gamestate
	Input       string
	LobbyCode   string // for lobby code creation or connection
}
