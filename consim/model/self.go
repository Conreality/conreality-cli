/* This is free and unencumbered software released into the public domain. */

package model

// Self
type Self struct{}

// IsPlayer
func (self *Self) IsPlayer() bool {
	return false
}

// IsRobot
func (self *Self) IsRobot() bool {
	return false
}
