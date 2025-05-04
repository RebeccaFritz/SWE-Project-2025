package main

import (
	"math/rand"
)

// initGameState initializes a gamestate struct
func initGameState() Gamestate {
	gs := Gamestate{
		Player1:     initPlayer(1),
		Player2:     initPlayer(2),
		Targets:     initTargets(),
		Projectiles: []Projectile{},
		Gameover:    false,
	}

	return gs
}

// initPlayer creates a blank player
func initPlayer(idx int) Player {
	Diameter := 100

	if idx == 1 {
		return Player{
			X: CANVAS_WIDTH / 2,
			Y: CANVAS_HEIGHT - Diameter /2,
			Diameter: Diameter,
			Velocity: 0,
			Idx: idx,
			Health: 5,
		}
	}

	return Player{
		X: CANVAS_WIDTH / 2,
		Y: Diameter / 2,
		Diameter: Diameter,
		Velocity: 0,
		Idx: idx,
		Health: 5,
	}
}

// initTargets creates the targets
func initTargets() []Target {
	targets := []Target{}

	for i := PLAYER_MOVE_LENGTH; i < CANVAS_WIDTH; i += PLAYER_MOVE_LENGTH {
		targets = append(targets, Target{X: i, Y: CANVAS_HEIGHT / 2, Velocity: 0, Diameter: 10, IsEnabled: true, Convert: rand.Intn(100)})
	}
	return targets
}

/* func initProjectiles() []Projectile {
	p1 := Projectile{350, 350, 10, 0, true, 0.5}
	p2 := Projectile{300, 300, 10, 0, true, 0.5}
	projectiles := []Projectile{p1, p2}
	return projectiles
}
*/
