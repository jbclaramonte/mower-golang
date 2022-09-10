package domain

type Mower struct {
	X           int
	Y           int
	Orientation Orientation // can be N, E, W, S
}

type Garden struct {
	Width  int
	Height int
}

type Orientation struct {
	Letter string
	Right  *Orientation
	Left   *Orientation
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
}
