package main

import "github.com/gorilla/websocket"

// maps websockets to clients
var CLIENTS = make(map[*websocket.Conn]*Client)

// the client struct
type Client struct {
	playerNum  int             // 1 or 2 (default 0)
	roomID     string          // (default "")
	username   string          // the client's username
	connection socket // the websocket connection this client is on
}
