package challenge09

import (
	"intcode"
	"testing"
)

func TestPuzzleInput(t *testing.T) {
	output := make(chan int64, 1)
	input := make(chan int64, 1)
	expectedVal := int64(2932210790)
	inputVal := int64(1)

	program := intcode.NewIntCodeProgram(puzzleInput, input, output, nil)

	input <- inputVal
	result, err := program.Execute()
	if err != nil {
		t.Errorf("Error running intcode: %v", err)
	}
	if result != expectedVal {
		t.Errorf("Unexpected result %d != %d", result, expectedVal)
	}

	outputVal := <-output
	if outputVal != expectedVal {
		t.Errorf("Unexpected output %d != %d", result, expectedVal)
	}
}

func TestPuzzleInputPart2(t *testing.T) {
	output := make(chan int64, 1)
	input := make(chan int64, 1)
	expectedVal := int64(73144)
	inputVal := int64(2)

	program := intcode.NewIntCodeProgram(puzzleInput, input, output, nil)

	input <- inputVal
	result, err := program.Execute()
	if err != nil {
		t.Errorf("Error running intcode: %v", err)
	}
	if result != expectedVal {
		t.Errorf("Unexpected result %d != %d", result, expectedVal)
	}

	outputVal := <-output
	if outputVal != expectedVal {
		t.Errorf("Unexpected output %d != %d", result, expectedVal)
	}
}
