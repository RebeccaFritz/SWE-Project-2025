package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// define a global array for the all the rooms
var ROOMS map[string]Room

// the Room struct contains all the information for one game
type Room struct {
	isFull      bool       // Room has two clients
	inGamestate bool       // is the room in the gamestate
	clients     [2]Client  // clients in the room
}

// runGameLoop is the entrypoint for a game session.
// printDebug controls whether the gamestate is printed to the screen
func runGameLoop(printDebug bool){
	gamestate := initGameState()

	for range time.Tick(TICK_DURATION){
		// input_queue := readPlayerInput() // this function should retrieve input that has been stored by the relevant pumps
		input_queue := []string{"move_left"} // for testing
		gamestate = updateGameState(gamestate, input_queue)
		// writeGameState() // player 2 needs to be flipped on read and write

		if(printDebug){
			log.Println("Gamestate")
			fmt.Printf("Projectiles: %+v\n", gamestate.projectiles)
			fmt.Printf("Targets: %+v\n", gamestate.targets)
			fmt.Printf("Player 1: %+v\n", gamestate.player1)
			fmt.Printf("Player 2: %+v\n\n", gamestate.player2)
		}
	}
}

// updateGameState adjusts the gamestate based on velocities and given player input
func updateGameState(gs Gamestate, input_queue []string)(Gamestate){
	gs = copyGameState(gs)

	gs.player1, gs.player2 = applyPlayerInputs(gs.player1, gs.player2, input_queue)
	updateProjectilePositions(gs.projectiles)
	updateTargetsPositions(gs.targets)
	handleProjectileTargetCollisions(gs.projectiles, gs.targets)

	return gs
}

// copyGameState returns a deep copy of the given game state
func copyGameState(gs Gamestate)(Gamestate){
	return gs
}

// applyPlayerInput adjusts the
func applyPlayerInputs(p1 Player, p2 Player, input_queue[]string)(Player, Player){
	for i:=range(input_queue){
		switch input_queue[i]{
			case "move_left":
				p1 = updatePlayerPosition(p1, "left")
				p2 = updatePlayerPosition(p2, "left")
			case "move_right":
				p1 = updatePlayerPosition(p1, "right")
				p2 = updatePlayerPosition(p2, "right")
			default:
				log.Printf("Client Input Error: unknown input '%s'\n", input_queue[i])
		}
	}

	return p1, p2
}

// updatePlayerPosition moves the given player the given direction based on the global PLAYER_MOVE_LENGTH
func updatePlayerPosition(p Player, direction string)(Player){
	if direction == "right"{
		p.x += PLAYER_MOVE_LENGTH
	} else if direction == "left" {
		p.x -= PLAYER_MOVE_LENGTH
	} else {
		log.Printf("Error: invalid move direction '%s'\n", direction)
	}

	return p
}

// updateTargetsPositions updates the positions of the targets, according to their velocity.
func updateTargetsPositions(targets []Target){
	for i:=range targets{
		if(targets[i].isEnabled){
			targets[i].y += targets[i].velocity
		}
	}
}

// updateProjectilesPositions updates the position of the projectiles, according to their velocity.
func updateProjectilePositions(projectiles []Projectile){
	for i:=range projectiles{
		if(projectiles[i].isEnabled){
			projectiles[i].y += projectiles[i].velocity
		}
	}
}

// handProjectileTargetCollisions checks for any collisions between the projectiles and the targets and applies the relevant velocity.
func handleProjectileTargetCollisions(projectiles []Projectile, targets []Target){
	for i := range targets{
      if(!targets[i].isEnabled) {
      	continue
      }

      for j := range projectiles{
         if(!projectiles[j].isEnabled){
         	continue
         }

         if(isColliding(targets[i], projectiles[j])){
         	targets[i].velocity += int(float64(projectiles[j].velocity) * projectiles[j].forceMultiplier)
            projectiles[j].isEnabled = false
            projectiles[j].velocity = 0
         }
      }
	}
}

// isColliding returns whether the given target and projectile are colliding.
func isColliding(target Target, projectile Projectile)(bool){
	displacement := distance(target.x, target.y, projectile.x, projectile.y)
   biggestDiameter := int(math.Max(float64(target.diameter), float64(projectile.diameter)))

   if(displacement <= biggestDiameter) {
    	return true
   }
   return false
}

func distance(x1 int, y1 int, x2 int, y2 int) int {
	return int(
		math.Sqrt(math.Pow(float64(x2) - float64(x1), 2) + math.Pow(float64(y2) - float64(y1), 2)))
}
