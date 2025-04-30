package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"time"
)

func deepCopyGamestate(gs Gamestate) Gamestate {
	copy := Gamestate{
		Player1:     gs.Player1,
		Player2:     gs.Player2,
		Projectiles: slices.Clone(gs.Projectiles),
		Targets:     slices.Clone(gs.Targets),
	}

	return copy
}

func reflectGamestate(oldGS Gamestate) Gamestate {
	gs := deepCopyGamestate(oldGS)

	gs.Player1.Y, gs.Player2.Y = gs.Player2.Y, gs.Player1.Y                             // vertical reflection
	gs.Player1.X, gs.Player2.X = CANVAS_HEIGHT-gs.Player1.X, CANVAS_HEIGHT-gs.Player2.X // horizontal reflection

	for j := range gs.Projectiles {
		gs.Projectiles[j].Y = CANVAS_HEIGHT - gs.Projectiles[j].Y // vertical reflection
		gs.Projectiles[j].X = CANVAS_HEIGHT - gs.Projectiles[j].X // horizontal reflection
	}

	for j := range gs.Targets {
		gs.Targets[j].Y = CANVAS_HEIGHT - gs.Targets[j].Y // vertical reflection
		gs.Targets[j].X = CANVAS_HEIGHT - gs.Targets[j].X // horizontal reflection
	}

	return gs
}

// runGameLoop updates the gamestate based on player input and writes it to the players in the room.
// printDebug controls whether the gamestate is printed to the console
func runGameLoop(printDebug bool, room *Room) {
	for range time.Tick(TICK_DURATION) {
		room.gamestate = updateGameState(room.gamestate, room.inputQueue)

		msg := msgStruct{MsgType: "gamestate", Gamestate: room.gamestate}
		handleWrite(1, msg, room.clients[0].connection)

		msg = msgStruct{MsgType: "gamestate", Gamestate: reflectGamestate(room.gamestate)}
		handleWrite(1, msg, room.clients[1].connection)

		// Clear the applied player input
		room.inputQueue = []InputQueueEntry{}

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
func updateGameState(gs Gamestate, input_queue []InputQueueEntry) Gamestate {
	gs = deepCopyGamestate(gs)

	gs.Player1, gs.Player2 = applyPlayerInputs(gs.Player1, gs.Player2, input_queue)
	updateProjectilePositions(gs.Projectiles)
	updateTargetsPositions(gs.Targets)
	handleProjectileTargetCollisions(gs.Projectiles, gs.Targets)
	gs.Player1.Health, gs.Player2.Health = handleTargetPlayerCollisions(gs.Targets, gs.Player1, gs.Player2)

	if gs.Player1.Health <= 0 || gs.Player2.Health <= 0 {
		gs.Gameover = true
	}

	return gs
}

// applyPlayerInput takes a input queue and applies it indiscriminately to the given players. see issue #85
func applyPlayerInputs(p1 Player, p2 Player, input_queue []InputQueueEntry) (Player, Player) {
	for i := range input_queue {
		switch input_queue[i].input {
		case "move_left", "move_right":
			if input_queue[i].player == 1 {
				p1 = updatePlayerPosition(p1, input_queue[i].input)
			} else {
				p2 = updatePlayerPosition(p2, input_queue[i].input)
			}
		case "launch_projectile":
			log.Println("Handle launching projectiles / base conversions here!")
		default:
			log.Printf("Client Input Error: unknown input '%s'\n", input_queue[i].input)
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

// handleProjectilePlayerCollisions updates any projectiles and players that are in collision conditions
func handleTargetPlayerCollisions(targets []Target, player1 Player, player2 Player) (int, int) {
	for i := range targets {
		if !targets[i].IsEnabled {
			continue
		}

		didReach, p1Health, p2Health := reachedOpponent(targets[i], player1, player2)
		player1.Health = p1Health
		player2.Health = p2Health

		if didReach {
			targets[i].IsEnabled = false
			targets[i].Velocity = 0
		}
	}

	return player1.Health, player2.Health
}

// reachedOpponent returns wheater the given target reached the opponent's collision zone
func reachedOpponent(target Target, player1 Player, player2 Player) (bool, int, int) {

	if (target.Y + target.Diameter) >= (CANVAS_HEIGHT - COLLISION_ZONE) {
		player1.Health -= 1
		return true, player1.Health, player2.Health
	} else if (target.Y - target.Diameter) <= COLLISION_ZONE {
		player2.Health -= 1
		return true, player1.Health, player2.Health
	}
	return false, player1.Health, player2.Health
}

func distance(x1 int, y1 int, x2 int, y2 int) int {
	return int(
		math.Sqrt(math.Pow(float64(x2)-float64(x1), 2) + math.Pow(float64(y2)-float64(y1), 2)))
}
