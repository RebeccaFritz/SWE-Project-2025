package main

type Gamestate struct {
	player1 Player;
	player2 Player;
	targets []Target;
	projectiles []Projectile;
	canvasWidth int;
	canvasHeight int;
}

type Player struct {
	x int;
	y int;
	diameter int;
	velocity int;
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
	forceMultiplier int;
}
