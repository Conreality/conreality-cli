/* This is free and unencumbered software released into the public domain. */

package model

// Team
type Team struct{}

// TeamPredicate
type TeamPredicate func(*Team) bool

// TeamPredicates
func TeamPredicates() map[string]TeamPredicate {
	return map[string]TeamPredicate{
		"is_alive": (*Team).IsAlive,
		"is_dead":  (*Team).IsDead,
	}
}

// IsPlayer
func (team *Team) IsAlive() bool {
	return true // TODO
}

// IsRobot
func (team *Team) IsDead() bool {
	return !team.IsAlive()
}
