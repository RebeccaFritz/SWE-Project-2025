package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// runGameLoop updates the gamestate based on player input and writes it to the players in the room.
// printDebug controls whether the gamestate is printed to the console
func runGameLoop(printDebug bool, room *Room) {
	for range time.Tick(TICK_DURATION) {
		room.gamestate = updateGameState(room.gamestate, room.inputQueue)

		gamestateMsg := msgStruct{MsgType: "gamestate", Gamestate: room.gamestate}
		handleWrite(1, gamestateMsg, room.clients[0].connection)
		handleWrite(1, gamestateMsg, room.clients[1].connection)

		// Clear the applied player input
		room.inputQueue = []string{}

		if printDebug {
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
func updateGameState(gs Gamestate, input_queue []string) Gamestate {
	gs.Player1, gs.Player2 = applyPlayerInputs(gs.Player1, gs.Player2, input_queue)
	updateProjectilePositions(gs.Projectiles)
	updateTargetsPositions(gs.Targets)
	handleProjectileTargetCollisions(gs.Projectiles, gs.Targets)

	return gs
}

// applyPlayerInput takes a input queue and applies it indiscriminately to the given players. see issue #85
func applyPlayerInputs(p1 Player, p2 Player, input_queue []string) (Player, Player) {
	for i := range input_queue {
		switch input_queue[i] {
		case "move_left":
			p1 = updatePlayerPosition(p1, input_queue[i])
			p2 = updatePlayerPosition(p2, input_queue[i])
		case "move_right":
			p1 = updatePlayerPosition(p1, input_queue[i])
			p2 = updatePlayerPosition(p2, input_queue[i])
		case "launch_projectile":
			log.Println("Handle launching projectiles / base conversions here!")
		default:
			log.Printf("Client Input Error: unknown input '%s'\n", input_queue[i])
		}
	}

	return p1, p2
}

// updatePlayerPosition moves the given player the given direction based on the global PLAYER_MOVE_LENGTH
func updatePlayerPosition(p Player, direction string) Player {
	if direction == "move_right" {
		p.X += PLAYER_MOVE_LENGTH
	} else if direction == "move_left" {
		p.X -= PLAYER_MOVE_LENGTH
	} else {
		log.Printf("Error: invalid move direction '%s'\n", direction)
	}

	return p
}

// updateTargetsPositions updates the positions of the targets, according to their velocity.
func updateTargetsPositions(targets []Target) {
	for i := range targets {
		if targets[i].IsEnabled {
			targets[i].Y += targets[i].Velocity
		}
	}
}

// updateProjectilesPositions updates the position of the projectiles, according to their velocity.
func updateProjectilePositions(projectiles []Projectile) {
	for i := range projectiles {
		if projectiles[i].IsEnabled {
			projectiles[i].Y += projectiles[i].Velocity
		}
	}
}

// handProjectileTargetCollisions checks for any collisions between the projectiles and the targets and applies the relevant velocity.
func handleProjectileTargetCollisions(projectiles []Projectile, targets []Target) {
	for i := range targets {
		if !targets[i].IsEnabled {
			continue
		}

		for j := range projectiles {
			if !projectiles[j].IsEnabled {
				continue
			}

			if isColliding(targets[i], projectiles[j]) {
				targets[i].Velocity += int(float64(projectiles[j].Velocity) * projectiles[j].ForceMult)
				projectiles[j].IsEnabled = false
				projectiles[j].Velocity = 0
			}
		}
	}
}

// isColliding returns whether the given target and projectile are colliding.
func isColliding(target Target, projectile Projectile) bool {
	displacement := distance(target.X, target.Y, projectile.X, projectile.Y)
	biggestDiameter := int(math.Max(float64(target.Diameter), float64(projectile.Diameter)))

	if displacement <= biggestDiameter {
		return true
	}
	return false
}

func distance(x1 int, y1 int, x2 int, y2 int) int {
	return int(
		math.Sqrt(math.Pow(float64(x2)-float64(x1), 2) + math.Pow(float64(y2)-float64(y1), 2)))
}
