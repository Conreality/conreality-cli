/* This is free and unencumbered software released into the public domain. */

package model

// Self
type Self struct{}

// SelfPredicate
type SelfPredicate func(*Self) bool

// SelfPredicates
func SelfPredicates() map[string]SelfPredicate {
	return map[string]SelfPredicate{
		"is_player":  (*Self).IsPlayer,
		"is_robot":   (*Self).IsRobot,
		"can_fly":    (*Self).CanFly,
		"can_move":   (*Self).CanMove,
		"has_legs":   (*Self).HasLegs,
		"has_wings":  (*Self).HasWings,
		"has_wheels": (*Self).HasWheels,
	}
}

// IsPlayer
func (self *Self) IsPlayer() bool {
	return false // TODO
}

// IsRobot
func (self *Self) IsRobot() bool {
	return false // TODO
}

// CanFly
func (self *Self) CanFly() bool {
	return false // TODO
}

// CanMove
func (self *Self) CanMove() bool {
	return false // TODO
}

// HasLegs
func (self *Self) HasLegs() bool {
	return false // TODO
}

// HasWings
func (self *Self) HasWings() bool {
	return false // TODO
}

// HasWheels
func (self *Self) HasWheels() bool {
	return false // TODO
}
