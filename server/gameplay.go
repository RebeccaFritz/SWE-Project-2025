package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func reflectGamestate(oldGS Gamestate) Gamestate {
	gs := deepCopyGamestate(oldGS)

	gs.Player1.Y, gs.Player2.Y = gs.Player2.Y, gs.Player1.Y                             // vertical reflection
	gs.Player1.X, gs.Player2.X = CANVAS_WIDTH-gs.Player1.X, CANVAS_WIDTH-gs.Player2.X // horizontal reflection

	for j := range gs.Projectiles {
		gs.Projectiles[j].Y = CANVAS_HEIGHT - gs.Projectiles[j].Y // vertical reflection
		gs.Projectiles[j].X = CANVAS_WIDTH - gs.Projectiles[j].X // horizontal reflection
	}

	for j := range gs.Targets {
		gs.Targets[j].Y = CANVAS_HEIGHT - gs.Targets[j].Y // vertical reflection
		gs.Targets[j].X = CANVAS_WIDTH - gs.Targets[j].X // horizontal reflection
	}

	return gs
}

// updateLeaderboard
func updateLeaderboard(gs Gamestate, room *Room) {
	switch {
	case gs.Player1.Health <= 0:
		if room.clients[1].username != "" {
			add_user(room.clients[1].username, DB)
			increment_wins(room.clients[1].username, DB)
		}
	case gs.Player2.Health <= 0:
		if room.clients[0].username != "" {
			add_user(room.clients[0].username, DB)
			increment_wins(room.clients[0].username, DB)
		}

	}
}

// runGameLoop updates the gamestate based on player input and writes it to the players in the room.
// printDebug controls whether the gamestate is printed to the console
func runGameLoop(printDebug bool, room *Room) {
	if printDebug{
		log.Println("Starting game with: ", room.clients[0].connection, " ",room.clients[1].connection)
	}

	for range time.Tick(TICK_DURATION) {
		room.gamestate = updateGameState(room.gamestate, room.inputQueue)

		msg := msgStruct{MsgType: "gamestate", Gamestate: room.gamestate}
		handleWrite(1, msg, &room.clients[0].connection)

		msg = msgStruct{MsgType: "gamestate", Gamestate: reflectGamestate(room.gamestate)}
		handleWrite(1, msg, &room.clients[1].connection)

		// Clear the applied player input
		room.inputQueue = []InputQueueEntry{}

		if room.gamestate.Gameover {
			updateLeaderboard(deepCopyGamestate(room.gamestate), room)
			LEADERBOARD = getLeaderboard(DB)

			handleWrite(1, LEADERBOARD, &room.clients[0].connection)
			handleWrite(1, LEADERBOARD, &room.clients[1].connection)
			break
		}

		if printDebug {
			log.Println("Gamestate")
			log.Printf("Projectiles: %+v\n", room.gamestate.Projectiles)
			log.Printf("Input queue: %+v\n", room.inputQueue)
			log.Printf("Targets: %+v\n", room.gamestate.Targets)
			log.Printf("Player 1: %+v\n", room.gamestate.Player1)
			log.Printf("Player 2: %+v\n\n", room.gamestate.Player2)
		}
	}

	delete(ROOMS, room.clients[0].roomID)
	log.Println("Room closed")
}

// updateGameState adjusts the gamestate based on velocities and given player input
func updateGameState(gs Gamestate, input_queue []InputQueueEntry) Gamestate {
	gs = deepCopyGamestate(gs)

	gs = applyPlayerInputs(deepCopyGamestate(gs), input_queue)
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
func applyPlayerInputs(gs Gamestate, input_queue []InputQueueEntry) Gamestate {
	for i := range input_queue {
		switch input_queue[i].input {
		case "move_left", "move_right":
			if input_queue[i].player == 1 {
				gs.Player1 = updatePlayerPosition(gs.Player1, input_queue[i].input, false)
			} else {
				gs.Player2 = updatePlayerPosition(gs.Player2, input_queue[i].input, true)
			}
		// case "launch_projectile":
		default:
			if input_queue[i].player == 1 {
				target := gs.Targets[i]
				for j := 0; j < len(gs.Targets); j++ {
					if gs.Targets[j].X == gs.Player1.X {
						target = gs.Targets[j]
					}
				}
				if doHexConversion(input_queue[i].input, target) {

					projectile := Projectile{gs.Player1.X, CANVAS_HEIGHT - 60, 10, -2, true, 1} // 60 pixels is the height of the player token showing in the Y direction
					gs.Projectiles = append(gs.Projectiles, projectile)
				}
			} else {
				target := gs.Targets[i]
				for j := 0; j < len(gs.Targets); j++ {
					if gs.Targets[j].X == gs.Player2.X {
						target = gs.Targets[j]
					}
				}
				if doHexConversion(input_queue[i].input, target) {
					projectile := Projectile{gs.Player2.X, 60, 10, 2, true, 1}
					gs.Projectiles = append(gs.Projectiles, projectile)
				}
			}
		}
	}
	return gs
}

// doHexConversion handles the game's hex->bin conversion mechanic to determine
// whether a projectile should be launched
func doHexConversion(input string, target Target) bool {
	log.Println("Input is:", input)
	convert := target.Convert
	binary := fmt.Sprintf("%08b", convert)
	log.Println("Target is:", binary)
	if input == binary {
		return true
	}
	return false
}

// updatePlayerPosition moves the given player the given direction based on the global PLAYER_MOVE_LENGTH
func updatePlayerPosition(p Player, direction string, isPlayer2 bool) Player {
	var P2mult int = 1
	var rightSide, leftSide int
	if isPlayer2 {
		P2mult = -1
		leftSide = CANVAS_WIDTH - (p.X + p.Diameter/2)  // reverse horizontal reflections
		rightSide = CANVAS_WIDTH - (p.X - p.Diameter/2) // reverse horizontal reflections
	} else {
		leftSide = (p.X - p.Diameter/2)
		rightSide = (p.X + p.Diameter/2)
	}

	//log.Println("move: (direction, p.X, p.left, p.right)", direction, p.X, leftSide, rightSide)

	if direction == "move_left" && leftSide > 0 {
		p.X -= PLAYER_MOVE_LENGTH * P2mult
	} else if direction == "move_right" && rightSide < CANVAS_WIDTH {
		p.X += PLAYER_MOVE_LENGTH * P2mult
	} else {
		if leftSide>0 || rightSide < CANVAS_WIDTH {
			log.Printf("Status: player tried to move out of bounds")
		} else {
			log.Printf("Error: invalid move direction '%s'\n", direction)
		}

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
				targets[i].Convert = rand.Intn(100)
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
