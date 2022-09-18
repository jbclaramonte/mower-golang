package service

import (
	"reflect"
	"testing"

	"github.com/jbclaramonte/mower-golang/core/domain"
)

// source: https://gist.github.com/samalba/6059502
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

var applyCommandTestCases = []struct {
	description   string
	garden        domain.Garden
	command       string
	initialState  domain.Mower
	expectedState domain.Mower
}{
	{
		description: "when mower is at the extrem north and oriented N it can't go forward",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "A",
		initialState: domain.Mower{
			X:           0,
			Y:           0,
			Orientation: domain.North,
		},
		expectedState: domain.Mower{
			X:           0,
			Y:           0,
			Orientation: domain.North,
		},
	},
	{
		description: "when mower is not on the edge and oriented N it can go forward",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "A",
		initialState: domain.Mower{
			X:           2,
			Y:           2,
			Orientation: domain.North,
		},
		expectedState: domain.Mower{
			X:           2,
			Y:           1,
			Orientation: domain.North,
		},
	},
	{
		description: "when mower is at the extrem east and oriented E it can't go forward",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "A",
		initialState: domain.Mower{
			X:           3,
			Y:           0,
			Orientation: domain.East,
		},
		expectedState: domain.Mower{
			X:           3,
			Y:           0,
			Orientation: domain.East,
		},
	},
	{
		description: "when mower is not on the edge and oriented E it can go forward",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "A",
		initialState: domain.Mower{
			X:           2,
			Y:           2,
			Orientation: domain.East,
		},
		expectedState: domain.Mower{
			X:           3,
			Y:           2,
			Orientation: domain.East,
		},
	},
	{
		description: "when mower is at the extrem west and oriented W it can't go forward",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "A",
		initialState: domain.Mower{
			X:           0,
			Y:           0,
			Orientation: domain.West,
		},
		expectedState: domain.Mower{
			X:           0,
			Y:           0,
			Orientation: domain.West,
		},
	},
	{
		description: "when mower is not on the edge and oriented W it can go forward",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "A",
		initialState: domain.Mower{
			X:           1,
			Y:           2,
			Orientation: domain.West,
		},
		expectedState: domain.Mower{
			X:           0,
			Y:           2,
			Orientation: domain.West,
		},
	},
	{
		description: "when mower is at the extrem south and oriented S it can't go forward",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "A",
		initialState: domain.Mower{
			X:           0,
			Y:           3,
			Orientation: domain.South,
		},
		expectedState: domain.Mower{
			X:           0,
			Y:           3,
			Orientation: domain.South,
		},
	},
	{
		description: "when mower is not on the edge and oriented S it can go forward",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "A",
		initialState: domain.Mower{
			X:           2,
			Y:           2,
			Orientation: domain.South,
		},
		expectedState: domain.Mower{
			X:           2,
			Y:           3,
			Orientation: domain.South,
		},
	},
	{
		description: "when mower turn right if its orientation was N then it will be E",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "D",
		initialState: domain.Mower{
			X:           2,
			Y:           2,
			Orientation: domain.North,
		},
		expectedState: domain.Mower{
			X:           2,
			Y:           2,
			Orientation: domain.East,
		},
	},
	{
		description: "when mower turn left if its orientation was N then it will be W",
		garden:      domain.Garden{Width: 4, Height: 4},
		command:     "G",
		initialState: domain.Mower{
			X:           2,
			Y:           2,
			Orientation: domain.North,
		},
		expectedState: domain.Mower{
			X:           2,
			Y:           2,
			Orientation: domain.West,
		},
	},
}

func TestService(t *testing.T) {
	for _, testCase := range applyCommandTestCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := ApplyCommand(testCase.garden, testCase.initialState, testCase.command)

			AssertEqual(t, result, testCase.expectedState)
		})
	}

	t.Run("When bad command is provided the app shoud panic", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("The code did not panic")
			}

		}()

		ApplyCommand(domain.Garden{Width: 4, Height: 4}, domain.Mower{}, "Q")

	})
}
