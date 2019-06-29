/* This is free and unencumbered software released into the public domain. */

package main

import "fmt"

// Headset
type Headset struct{}

// Speak
func (headset *Headset) Speak(message string) {
	fmt.Printf(">>> %s\n", message)
}

// Agent
type Agent struct{}

// IsPlayer
func (agent *Agent) IsPlayer() bool {
	return false // TODO
}

// IsRobot
func (agent *Agent) IsRobot() bool {
	return false // TODO
}

// CanFly
func (agent *Agent) CanFly() bool {
	return false // TODO
}

// CanMove
func (agent *Agent) CanMove() bool {
	return false // TODO
}

// HasLegs
func (agent *Agent) HasLegs() bool {
	return false // TODO
}

// HasWings
func (agent *Agent) HasWings() bool {
	return false // TODO
}

// HasWheels
func (agent *Agent) HasWheels() bool {
	return false // TODO
}

// Name
func (agent *Agent) Name() string {
	return "Unknown" // TODO
}

// Game
type Game struct{}

// IsPaused
func (game *Game) IsPaused() bool {
	return false // TODO
}

// IsPrivate
func (game *Game) IsPrivate() bool {
	return false // TODO
}

// IsPublic
func (game *Game) IsPublic() bool {
	return !game.IsPrivate()
}

// IsStarted
func (game *Game) IsStarted() bool {
	return false // TODO
}

// IsStopped
func (game *Game) IsStopped() bool {
	return false // TODO
}

// Theater
type Theater struct{}

// Unit
type Unit struct{}

// IsAlive
func (unit *Unit) IsAlive() bool {
	return true // TODO
}

// IsDead
func (unit *Unit) IsDead() bool {
	return !unit.IsAlive()
}
