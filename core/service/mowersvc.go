package service

import (
	"mower/core/domain"
)

type Command interface {
	apply(mower domain.Mower) domain.Mower
}

type forward struct {
}

func (c forward) apply(mower domain.Mower) domain.Mower {
	var result domain.Mower
	if mower.Orientation == "N" {
		if mower.Y-1 < 0 {
			result = mower
		} else {
			result = domain.Mower{
				X:           mower.X,
				Y:           mower.Y - 1,
				Orientation: mower.Orientation,
			}
		}
	}

	return result
}

func ApplyComand(mower domain.Mower, command string) domain.Mower {
	var cmd Command
	if command == "N" {
		cmd = forward{}
	}
	return cmd.apply(mower)
}
