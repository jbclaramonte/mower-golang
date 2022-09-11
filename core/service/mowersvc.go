package service

import (
	"github.com/jbclaramonte/mower-golang/core/domain"
)

type Command interface {
	apply(garden domain.Garden, mower domain.Mower) domain.Mower
}

type forward struct {
}

func (c forward) apply(garden domain.Garden, mower domain.Mower) domain.Mower {
	var result domain.Mower
	if mower.Orientation == domain.North {
		if mower.Y-1 < 0 {
			result = mower
		} else {
			result = domain.Mower{
				X:           mower.X,
				Y:           mower.Y - 1,
				Orientation: mower.Orientation,
			}
		}
	} else if mower.Orientation == domain.East {
		if mower.X+1 < garden.Width {
			result = domain.Mower{
				X:           mower.X + 1,
				Y:           mower.Y,
				Orientation: mower.Orientation,
			}
		} else {
			result = mower
		}
	} else if mower.Orientation == domain.West {
		if mower.X-1 > 0 {
			result = domain.Mower{
				X:           mower.X - 1,
				Y:           mower.Y,
				Orientation: mower.Orientation,
			}
		} else {
			result = mower
		}
	} else if mower.Orientation == domain.South {
		if mower.Y+1 < garden.Height {
			result = domain.Mower{
				X:           mower.X,
				Y:           mower.Y + 1,
				Orientation: mower.Orientation,
			}
		} else {
			result = mower
		}
	}

	return result
}

func createCommand(command string) Command {
	if command == "A" {
		return forward{}
	} else if command == "D" {
		return turnRight{}
	} else if command == "L" {
		return turnLeft{}
	}
	// panic("Unknown command" + command)
	println("------------------------")
	return nil
}

func ApplyComand(garden domain.Garden, mower domain.Mower, command string) domain.Mower {
	cmd := createCommand(command)

	return cmd.apply(garden, mower)
}

type turnRight struct{}

func (r turnRight) apply(garden domain.Garden, mower domain.Mower) domain.Mower {
	mower.Orientation = *mower.Orientation.Right
	return mower
}

type turnLeft struct{}

func (l turnLeft) apply(garden domain.Garden, mower domain.Mower) domain.Mower {
	mower.Orientation = *mower.Orientation.Left
	return mower
}
