package main

import (
	"math"
)

// define a global array for the all the rooms
var ROOMS map[string]Room

// the Room struct contains all the information for one game
type Room struct {
	isFull      bool       // Room has two clients
	inGamestate bool       // is the room in the gamestate
	clients     [2]Client  // clients in the room
	targets     [10]Target // struct containing information for each first target
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
            targets[i].velocity += projectiles[j].velocity * projectiles[j].forceMultiplier
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
