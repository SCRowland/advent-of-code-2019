package main

import (
	"sliceutils"
	"testing"
)

func TestIterateInstructions(t *testing.T) {
	instructions := []int{1, 2, 3, 4}
	got := NewIntCodeProgram(instructions)

	if got.programCounter != 0 {
		t.Errorf("NewIntCodeProgram(%v).programCounter = %d, should be 0", instructions, got)
	}

	if !sliceutils.TestIntSliceEq32(got.rawInstructions, instructions) {
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

	got := NewIntCodeProgram(instructions)
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
		got := NewIntCodeProgram(tt.programText)
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
		got := NewIntCodeProgram(tt.programText)
		err := got.Execute()
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
		got := NewIntCodeProgram(tt.programText)
		err := got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !sliceutils.TestIntSliceEq32(got.rawInstructions, tt.expectedFinalState) {
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
		got := NewIntCodeProgram(tt.programText)
		err := got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !sliceutils.TestIntSliceEq32(got.rawInstructions, tt.expectedFinalState) {
			t.Errorf("NewIntCodeProgram(%v) ended %v, not %v", tt.programText, got.rawInstructions, tt.expectedFinalState)
		}
	}
}

func TestInputOperation(t *testing.T) {
	// TODO!
	// how to mock input? - encapsulate inout to a func and mock that I guess?
}
func TestOutputOperation(t *testing.T) {
	// TODO
	// how to capture output for testing?
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
		got := NewIntCodeProgram(tt.programText)
		err := got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !sliceutils.TestIntSliceEq32(got.rawInstructions, tt.expectedFinalState) {
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

var instructions = []struct {
	instruction        int
	expectedOpcode     int
	expectedModeParam1 int
	expectedModeParam2 int
	expectedModeParam3 int
}{
	{
		1002,
		2,
		0,
		1,
		0,
	},
	{
		11102,
		2,
		1,
		1,
		1,
	},
	{
		10105,
		5,
		1,
		0,
		1,
	},
	{
		2,
		2,
		0,
		0,
		0,
	},
}

func TestInstructionParsing(t *testing.T) {
	for _, testData := range instructions {
		opcode, paramMode1, paramMode2, paramMode3 := parseInstruction(testData.instruction)
		if opcode != testData.expectedOpcode {
			t.Errorf("parseInstruction(%d) = %d, _, _, _, should be %d, _, _, _",
				testData.instruction,
				opcode,
				testData.expectedOpcode,
			)
		}
		if paramMode1 != testData.expectedModeParam1 {
			t.Errorf("parseInstruction(%d) = _, %d, _, _, should be _, %d, _, _",
				testData.instruction,
				paramMode1,
				testData.expectedModeParam1,
			)
		}
		if paramMode2 != testData.expectedModeParam2 {
			t.Errorf("parseInstruction(%d) = _, _, %d, _, should be _, _, %d, _",
				testData.instruction,
				paramMode2,
				testData.expectedModeParam2,
			)
		}
		if paramMode3 != testData.expectedModeParam3 {
			t.Errorf("parseInstruction(%d) = _, _, _, %d should be _, _, _, %d",
				testData.instruction,
				paramMode3,
				testData.expectedModeParam3,
			)
		}
	}
}

var fetchValues = []struct {
	instructionText []int
	paramValue      int
	paramMode       int
	expectedValue   int
}{
	{
		[]int{0, 3, 2, 5},
		1,
		positional,
		3,
	},
	{
		[]int{0, 3, 2, 5},
		1,
		immediate,
		1,
	},
	{
		[]int{0, 3, 4, 5, 6},
		2,
		positional,
		4,
	},
	{
		[]int{0, 3, 4, 5, 6},
		2,
		immediate,
		2,
	},
}

func TestFetchValue(t *testing.T) {
	for _, testData := range fetchValues {
		program := NewIntCodeProgram(testData.instructionText)
		got, err := program.fetchValue(testData.paramValue, testData.paramMode)
		if err != nil {

		}
		if got != testData.expectedValue {
			t.Errorf("NewIntCodeProgram(%v).fetchValue(%d, %d) = %d, not %d",
				testData.instructionText,
				testData.paramValue,
				testData.paramMode,
				got,
				testData.expectedValue,
			)
		}
	}
}

func TestGetResult(t *testing.T) {
	var programIntructions = []int{1, 2, 3, 4}
	program := NewIntCodeProgram(programIntructions)
	got := program.GetResult()
	if got != 1 {
		t.Errorf("NewIntCodeProgram(%v).GetResult = %d, not 1", programIntructions, got)
	}
}

func TestSetInitialError(t *testing.T) {
	var programIntructions = []int{1, 2, 3, 4}
	program := NewIntCodeProgram(programIntructions)
	program.SetInitialError(123, 456)
	if program.rawInstructions[1] != 123 {
		t.Errorf("program.rawInstructions[1] = %d not 123", program.rawInstructions[1])
	}
	program.SetInitialError(123, 456)
	if program.rawInstructions[2] != 456 {
		t.Errorf("program.rawInstructions[2] = %d not 456", program.rawInstructions[1])
	}
}
