package main

import("slices")

type Gamestate struct {
	Player1     Player
	Player2     Player
	Targets     []Target
	Projectiles []Projectile
	Gameover    bool
}

func deepCopyGamestate(gs Gamestate) Gamestate {
	copy := Gamestate{
		Player1:     gs.Player1,
		Player2:     gs.Player2,
		Projectiles: slices.Clone(gs.Projectiles),
		Targets:     slices.Clone(gs.Targets),
		Gameover: 	 gs.Gameover,
	}

	return copy
}

type Player struct {
	X        int
	Y        int
	Diameter int
	Velocity int
	Idx      int
	Health   int
}

type Target struct {
	X         int
	Y         int
	Diameter  int
	Velocity  int
	IsEnabled bool
	Convert   int
}

type Projectile struct {
	X         int
	Y         int
	Diameter  int
	Velocity  int
	IsEnabled bool
	ForceMult float64
}
