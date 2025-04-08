package main
import(
	"time"
)

/*
CONSTANTS THAT DEFINE ALL GAME STATES *USE THEM CAREFULLY*
*/

// this is the number of targets structures, not the current number of enabled targets.
const NUM_TARGETS int = 3
const CANVAS_WIDTH = 400;
const CANVAS_HEIGHT = 400;
const PLAYER_MOVE_LENGTH = 50;
const TICK_DURATION = 1 * time.Second

type Gamestate struct {
	player1 Player;
	player2 Player;
	targets []Target;
	projectiles []Projectile;
}

type Player struct {
	x int;
	y int;
	diameter int;
	velocity int;
	idx int;
}

type Target struct {
	x int;
	y int;
	diameter int;
	velocity int;
	isEnabled bool;
}

type Projectile struct {
	x int;
	y int;
	diameter int;
	velocity int;
	isEnabled bool;
	forceMultiplier float64;
}
