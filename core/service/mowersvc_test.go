package service

import (
	"mower/core/domain"
	"reflect"
	"testing"
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
	command       string
	initialState  domain.Mower
	expectedState domain.Mower
}{
	{
		description: "when mower is at the extrem top it can't go forward",
		command:     "N",
		initialState: domain.Mower{
			X:           0,
			Y:           0,
			Orientation: "N",
		},
		expectedState: domain.Mower{
			X:           0,
			Y:           0,
			Orientation: "N",
		},
	},
}

func TestService(t *testing.T) {
	for _, testCase := range applyCommandTestCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := ApplyComand(testCase.initialState, testCase.command)

			AssertEqual(t, result, testCase.expectedState)
		})
	}
}
