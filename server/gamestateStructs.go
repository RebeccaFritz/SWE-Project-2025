package main

type Gamestate struct {
	Player1     Player
	Player2     Player
	Targets     []Target
	Projectiles []Projectile
}

type Player struct {
	X        int
	Y        int
	Diameter int
	Velocity int
	Idx      int
}

type Target struct {
	X         int
	Y         int
	Diameter  int
	Velocity  int
	IsEnabled bool
}

type Projectile struct {
	X         int
	Y         int
	Diameter  int
	Velocity  int
	IsEnabled bool
	ForceMult float64
}
