package main

import (
	"time"
)

/*
CONSTANTS THAT DEFINE ALL GAME STATES *USE THEM CAREFULLY*
*/
const CANVAS_WIDTH = 400
const CANVAS_HEIGHT = 600
const PLAYER_MOVE_LENGTH = 50
const TICK_DURATION = time.Second / 30
const COLLISION_ZONE = 50 // this is equal to the radius of the player token
