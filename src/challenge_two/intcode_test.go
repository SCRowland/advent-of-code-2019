package main

import "testing"

func TestIterateInstructions(t *testing.T) {
	instructions := []int{1, 2, 3, 4}
	got, err := NewIntCodeProgram(instructions)
	if err != nil {
		t.Errorf("NewIntCodeProgram(%v) error: %s", instructions, err)
	}

	if got.programCounter != 0 {
		t.Errorf("NewIntCodeProgram(%v).programCounter = %d, should be 0", instructions, got)
	}

	if !TestIntSliceEq(got.rawInstructions, instructions) {
		t.Errorf("NewIntCodeProgram(%v).rawInstructions = %v, should be %v", instructions, got.rawInstructions, instructions)
	}

	if len(got.rawInstructions) != 4 {
		t.Errorf("len(NewIntCodeProgram(%v).instructions) = %d, should be %d", instructions, len(got.rawInstructions), 4)
	}
}

func TestExectuteNext(t *testing.T) {
	instructions := []int{
		1, 0, 0, 0,
		2, 0, 0, 0,
		1, 0, 0, 0,
		99,
	}

	got, err := NewIntCodeProgram(instructions)
	if err != nil {
		t.Errorf("NewIntCodeProgram(%v) error: %s", instructions, err)
	}

	got.Execute()

	if got.programCounter != 12 {
		t.Errorf("NewIntCodeProgram(%v).programCounter = %d, should be 4", instructions, got.programCounter)
	}
}

var haltPrograms = []struct {
	programText            []int
	expectedProgramCounter int
}{
	{[]int{
		1, 0, 0, 0,
		2, 0, 0, 0,
		1, 0, 0, 0,
		99, 0, 0, 0,
		1, 0, 0, 0,
	}, 12},
	{[]int{
		1, 0, 0, 0,
		2, 0, 0, 0,
		1, 0, 0, 0,
		2, 0, 0, 0,
		1, 0, 0, 0,
		2, 0, 0, 0,
		99,
	}, 24},
	{[]int{
		99,
	}, 0},
}

func TestExecuteHalt(t *testing.T) {
	for _, tt := range haltPrograms {
		got, err := NewIntCodeProgram(tt.programText)
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		got.Execute()

		if got.programCounter != tt.expectedProgramCounter {
			t.Errorf("NewIntCodeProgram(%v).programCounter = %d, should be %d", tt.programText, got.programCounter, tt.expectedProgramCounter)
		}
	}
}

var badOpcodePrograms = []struct {
	programText []int
}{
	{[]int{
		3, 0, 0, 0,
	}},
	{[]int{
		-1, 0, 0, 0,
	}},
	{[]int{
		98, 0, 0, 0,
	}},
}

func TestBadOpcode(t *testing.T) {
	for _, tt := range badOpcodePrograms {
		got, err := NewIntCodeProgram(tt.programText)

		err = got.Execute()
		if err == nil {
			t.Errorf("NewIntCodeProgram(%v) opCode error not thrown", tt.programText)
		}
	}
}

var addOpcodePrograms = []struct {
	programText        []int
	expectedFinalState []int
}{
	{
		[]int{
			1, 5, 6, 7,
			99, 25, 14, 0,
		},
		[]int{
			1, 5, 6, 7,
			99, 25, 14, 39,
		},
	},
}

func TestAdditionOpcode(t *testing.T) {
	for _, tt := range addOpcodePrograms {
		got, err := NewIntCodeProgram(tt.programText)

		err = got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !TestIntSliceEq(got.rawInstructions, tt.expectedFinalState) {
			t.Errorf("NewIntCodeProgram(%v) ended %v, not %v", tt.programText, got.rawInstructions, tt.expectedFinalState)
		}
	}
}

var multiplyOpcodePrograms = []struct {
	programText        []int
	expectedFinalState []int
}{
	{
		[]int{
			2, 5, 6, 7,
			99, 25, 14, 0,
		},
		[]int{
			2, 5, 6, 7,
			99, 25, 14, 350,
		},
	},
}

func TestMultiplyOpcode(t *testing.T) {
	for _, tt := range multiplyOpcodePrograms {
		got, err := NewIntCodeProgram(tt.programText)

		err = got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !TestIntSliceEq(got.rawInstructions, tt.expectedFinalState) {
			t.Errorf("NewIntCodeProgram(%v) ended %v, not %v", tt.programText, got.rawInstructions, tt.expectedFinalState)
		}
	}
}

var complexPrograms = []struct {
	programText        []int
	expectedFinalState []int
}{
	{
		[]int{
			1, 9, 10, 11,
			1, 11, 10, 9,
			99, 3, 4, 0,
		},
		[]int{
			1, 9, 10, 11,
			1, 11, 10, 9,
			99, 11, 4, 7,
		},
	},
}

func TestComplexProgram(t *testing.T) {
	for _, tt := range complexPrograms {
		got, err := NewIntCodeProgram(tt.programText)

		err = got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !TestIntSliceEq(got.rawInstructions, tt.expectedFinalState) {
			t.Errorf("NewIntCodeProgram(%v) ended %v, not %v", tt.programText, got.rawInstructions, tt.expectedFinalState)
		}
	}
}

var inputProgram = []int{
	1, 0, 0, 3,
	1, 1, 2, 3,
	1, 3, 4, 3,
	1, 5, 0, 3,
	2, 1, 10, 19,
	1, 19, 5, 23,
	2, 23, 6, 27,
	1, 27, 5, 31,
	2, 6, 31, 35,
	1, 5, 35, 39,
	2, 39, 9, 43,
	1, 43, 5, 47,
	1, 10, 47, 51,
	1, 51, 6, 55,
	1, 55, 10, 59,
	1, 59, 6, 63,
	2, 13, 63, 67,
	1, 9, 67, 71,
	2, 6, 71, 75,
	1, 5, 75, 79,
	1, 9, 79, 83,
	2, 6, 83, 87,
	1, 5, 87, 91,
	2, 6, 91, 95,
	2, 95, 9, 99,
	1, 99, 6, 103,
	1, 103, 13, 107,
	2, 13, 107, 111,
	2, 111, 10, 115,
	1, 115, 6, 119,
	1, 6, 119, 123,
	2, 6, 123, 127,
	1, 127, 5, 131,
	2, 131, 6, 135,
	1, 135, 2, 139,
	1, 139, 9, 0,
	99, 2, 14, 0,
	0,
}

func TestChallengeInput(t *testing.T) {
	expectedResult := 5866663

	got, err := NewIntCodeProgram(inputProgram)

	got.SetInitialError(12, 2)
	err = got.Execute()
	if err != nil {
		t.Errorf("NewIntCodeProgram(challengeProgram) error: %s", err)
	}

	result := got.GetResult()
	if result != expectedResult {
		t.Errorf("NewIntCodeProgram(challengeProgram) = %d", result)
	}
}
