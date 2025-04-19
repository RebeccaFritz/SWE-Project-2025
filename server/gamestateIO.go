package main

import (
	"github.com/gorilla/websocket"
)

func writeGameState(socket *websocket.Conn, gs Gamestate) {
	gamestate_msg := msgStruct { MsgType: "gamestate", Gamestate: gs }
	handleWrite(1, gamestate_msg, socket)

	// _, msgStruct, _ := handleRead(socket)

	// if msgStruct.MsgType != "pong"{
	// 	log.Println("Error: did not recieve gamestate pong from client!")

	// } else {
	// 	log.Printf("Gamestate sent succesfully!")
	// }
}


// Implement the logic to read the game state from a file or database
// // Ensure that player 2's data is flipped after reading
