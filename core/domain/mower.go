package domain

import (
	"fmt"
)

type Mower struct {
	X           int
	Y           int
	Orientation Orientation // can be N, E, W, S
}

type Garden struct {
	Width  int
	Height int
}

func GetOrientation(orientation string) (*Orientation, error) {
	switch orientation {
	case "N":
		return &North, nil
	case "S":
		return &South, nil
	case "E":
		return &East, nil
	case "W":
		return &West, nil
	}
	return nil, fmt.Errorf("no orientation for %v", orientation)
}

type Orientation struct {
	Letter string
	Right  *Orientation `json:"-"`
	Left   *Orientation `json:"-"`
}

var North = Orientation{
	Letter: "N",
}

var East = Orientation{
	Letter: "E",
}

var West = Orientation{
	Letter: "W",
}

var South = Orientation{
	Letter: "S",
}

func init() {
	North.Right = &East
	East.Left = &North
	East.Right = &South
	South.Left = &East
	South.Right = &West
	West.Left = &South
	West.Right = &North
	North.Left = &West
}
