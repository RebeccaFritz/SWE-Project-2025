package main

import (
	"fmt"
	"log"
	"math"
	"time"
)


// the Room struct contains all the information for one game
type Room struct {
	gamestate	Gamestate
	inputQueue	[]string
	clients     [2]*Client  // clients in the room
}

// flipGameState flips the given gamestate so that player 2 can render it correctly
func flipGameState(gs Gamestate) Gamestate {
	gs = copyGameState(gs)

	return gs
}

// runGameLoop updates the gamestate based on player input and writes it to the players in the room
// printDebug controls whether the gamestate is printed to the screen
func runGameLoop(printDebug bool, room *Room){
	// log.Println("I am the game thread and here is the room i got: ", *room)
	for range time.Tick(TICK_DURATION){
		room.gamestate = updateGameState(room.gamestate, room.inputQueue)

		gamestateMsg := msgStruct { MsgType: "gamestate", Gamestate: room.gamestate }
		handleWrite(1, gamestateMsg, room.clients[0].connection)
		handleWrite(1, gamestateMsg, room.clients[1].connection)

		// Clear applied player input
		room.inputQueue = []string{}

		if(printDebug){
			log.Println("Gamestate")
			fmt.Printf("Projectiles: %+v\n", room.gamestate.Projectiles)
			fmt.Printf("Input queue: %+v\n", room.inputQueue)
			fmt.Printf("Targets: %+v\n", room.gamestate.Targets)
			fmt.Printf("Player 1: %+v\n", room.gamestate.Player1)
			fmt.Printf("Player 2: %+v\n\n", room.gamestate.Player2)
		}
	}
}

// updateGameState adjusts the gamestate based on velocities and given player input
func updateGameState(gs Gamestate, input_queue []string)(Gamestate){
	gs = copyGameState(gs)

	gs.Player1, gs.Player2 = applyPlayerInputs(gs.Player1, gs.Player2, input_queue)
	updateProjectilePositions(gs.Projectiles)
	updateTargetsPositions(gs.Targets)
	handleProjectileTargetCollisions(gs.Projectiles, gs.Targets)

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
		p.X += PLAYER_MOVE_LENGTH
	} else if direction == "left" {
		p.X -= PLAYER_MOVE_LENGTH
	} else {
		log.Printf("Error: invalid move direction '%s'\n", direction)
	}

	return p
}

// updateTargetsPositions updates the positions of the targets, according to their velocity.
func updateTargetsPositions(targets []Target){
	for i:=range targets{
		if(targets[i].IsEnabled){
			targets[i].Y += targets[i].Velocity
		}
	}
}

// updateProjectilesPositions updates the position of the projectiles, according to their velocity.
func updateProjectilePositions(projectiles []Projectile){
	for i:=range projectiles{
		if(projectiles[i].IsEnabled){
			projectiles[i].Y += projectiles[i].Velocity
		}
	}
}

// handProjectileTargetCollisions checks for any collisions between the projectiles and the targets and applies the relevant velocity.
func handleProjectileTargetCollisions(projectiles []Projectile, targets []Target){
	for i := range targets{
      if(!targets[i].IsEnabled) {
      	continue
      }

      for j := range projectiles{
         if(!projectiles[j].IsEnabled){
         	continue
         }

         if(isColliding(targets[i], projectiles[j])){
         	targets[i].Velocity += int(float64(projectiles[j].Velocity) * projectiles[j].ForceMult)
            projectiles[j].IsEnabled = false
            projectiles[j].Velocity = 0
         }
      }
	}
}

// isColliding returns whether the given target and projectile are colliding.
func isColliding(target Target, projectile Projectile)(bool){
	displacement := distance(target.X, target.Y, projectile.X, projectile.Y)
   biggestDiameter := int(math.Max(float64(target.Diameter), float64(projectile.Diameter)))

   if(displacement <= biggestDiameter) {
    	return true
   }
   return false
}

func distance(x1 int, y1 int, x2 int, y2 int) int {
	return int(
		math.Sqrt(math.Pow(float64(x2) - float64(x1), 2) + math.Pow(float64(y2) - float64(y1), 2)))
}
