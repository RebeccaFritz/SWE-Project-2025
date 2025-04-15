package main

// initGameState initializes a gamestate struct
func initGameState()(Gamestate){
	gs := Gamestate {
		player1: initPlayer(1),
		player2: initPlayer(2),
		targets: initTargets(),
		projectiles: initProjectiles(),
	}

	return gs
}

// initPlayer creates a blank player
func initPlayer(idx int)(Player){
	return Player { x: 0, y: 0, diameter: 0, velocity: 0, idx: idx}
}

// initTargets creates the targets
func initTargets()([]Target){
	t1 := Target{350, 350, 10, 0, true}
	t2 := Target{300, 300, 10, 0, true}
	targets := []Target{t1, t2}
	return targets
}

func initProjectiles()([]Projectile){
	p1 := Projectile{350, 350, 10, 0, true, 0.5}
	p2 := Projectile{300, 300, 10, 0, true, 0.5}
	projectiles := []Projectile{p1, p2}
	return projectiles
}
