package main

// define a global array for the all the rooms
var ROOMS map[string]Room

// the Room struct contains all the information for one game
type Room struct {
	isFull      bool       // Room has two clients
	inGamestate bool       // is the room in the gamestate
	clients     [2]Client  // clients in the room
	targets     [10]Target // struct containing information for each first target
}

// the Target struct
type Target struct {
	twosComp   int    // two's complement number
	baseTen    int    // base 10 number
	hasBoost   bool   // does this target have a boost
	isOnScreen bool   // is this target on screen
	position0  [2]int // position as the Target would appear on player 0's screen
	position1  [2]int // position as the Target would appear on player 1's screen
}
