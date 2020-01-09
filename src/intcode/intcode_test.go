package intcode

import (
	"sliceutils"
	"testing"
)

func TestIterateInstructions(t *testing.T) {
	instructions := []int64{1, 2, 3, 4}
	got := NewIntCodeProgram(instructions, nil, nil, nil)

	if got.programCounter != 0 {
		t.Errorf("NewIntCodeProgram(%v).programCounter = %d, should be 0", instructions, got)
	}

	if !sliceutils.TestIntSliceEq(got.rawInstructions, instructions) {
		t.Errorf("NewIntCodeProgram(%v).rawInstructions = %v, should be %v", instructions, got.rawInstructions, instructions)
	}

	if len(got.rawInstructions) != 4 {
		t.Errorf("len(NewIntCodeProgram(%v).instructions) = %d, should be %d", instructions, len(got.rawInstructions), 4)
	}
}

func TestNoHaltInstruction(t *testing.T) {
	instructions := []int64{1, 2, 3, 4, 13}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err == nil {
		t.Errorf("Error not thrown without halt instruction")
	}
}

func TestJumpIfTrueTrue(t *testing.T) {
	instructions := []int64{1105, 3, 11, 99, 3, 99, 3, 2, 3, 2, 3, 99}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("jumpIfTrue returned err %v", err)
	}

	if program.programCounter != 11 {
		t.Errorf("jumpIfTrue stored %d, not %d", program.programCounter, 11)
	}
}

func TestJumpIfTrueFalse(t *testing.T) {
	instructions := []int64{1105, 0, 11, 99, 3, 99, 3, 2, 3, 2, 3, 99}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("jumpIfTrue returned err %v", err)
	}

	if program.programCounter != 3 {
		t.Errorf("jumpIfTrue stored %d, not %d", program.programCounter, 3)
	}
}

func TestJumpIfFalseTrue(t *testing.T) {
	instructions := []int64{1106, 0, 11, 99, 3, 99, 3, 2, 3, 2, 3, 99}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("jumpIfFalse returned err %v", err)
	}

	if program.programCounter != 11 {
		t.Errorf("jumpIfFalse stored %d, not %d", program.programCounter, 11)
	}
}

func TestJumpIfFalseFalse(t *testing.T) {
	instructions := []int64{1106, 1, 11, 99, 3, 99, 3, 2, 3, 2, 3, 99}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("jumpIfFalse returned err %v", err)
	}

	if program.programCounter != 3 {
		t.Errorf("jumpIfFalse stored %d, not %d", program.programCounter, 3)
	}
}

func TestEqualsTrue(t *testing.T) {
	instructions := []int64{1108, 73, 73, 0, 99}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("equals returned err %v", err)
	}

	if program.rawInstructions[0] != 1 {
		t.Errorf("equals stored %d, not %d", program.rawInstructions[0], 1)
	}
}

func TestEqualsFalse(t *testing.T) {
	instructions := []int64{1108, 73, -73, 0, 99}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("equals returned err %v", err)
	}

	if program.rawInstructions[0] != 0 {
		t.Errorf("equals stored %d, not %d", program.rawInstructions[0], 0)
	}
}

func TestJumpIfFalse(t *testing.T) {

}

func TestLessThanTrue(t *testing.T) {
	instructions := []int64{1107, 2, 3, 0, 99}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("LessThanProgram returned err %v", err)
	}

	if program.rawInstructions[0] != 1 {
		t.Errorf("Less Than Instruction stored %d, not %d", program.rawInstructions[0], 1)
	}
}

func TestLessThanFalse(t *testing.T) {
	instructions := []int64{1107, 3, 2, 0, 99}
	program := NewIntCodeProgram(instructions, nil, nil, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("LessThanProgram returned err %v", err)
	}

	if program.rawInstructions[0] != 0 {
		t.Errorf("Less Than Instruction stored %d, not %d", program.rawInstructions[0], 1)
	}
}

func TestEquals(t *testing.T) {

}

func TestExectuteNext(t *testing.T) {
	instructions := []int64{
		1, 0, 0, 0,
		2, 0, 0, 0,
		1, 0, 0, 0,
		99,
	}

	got := NewIntCodeProgram(instructions, nil, nil, nil)
	got.Execute()

	if got.programCounter != int64(12) {
		t.Errorf("NewIntCodeProgram(%v).programCounter = %d, should be 4", instructions, got.programCounter)
	}
}

var haltPrograms = []struct {
	programText            []int64
	expectedProgramCounter int64
}{
	{[]int64{
		1, 0, 0, 0,
		2, 0, 0, 0,
		1, 0, 0, 0,
		99, 0, 0, 0,
		1, 0, 0, 0,
	}, 12},
	{[]int64{
		1, 0, 0, 0,
		2, 0, 0, 0,
		1, 0, 0, 0,
		2, 0, 0, 0,
		1, 0, 0, 0,
		2, 0, 0, 0,
		99,
	}, 24},
	{[]int64{
		99,
	}, 0},
}

func TestExecuteHalt(t *testing.T) {
	for _, tt := range haltPrograms {
		got := NewIntCodeProgram(tt.programText, nil, nil, nil)
		got.Execute()

		if got.programCounter != tt.expectedProgramCounter {
			t.Errorf("NewIntCodeProgram(%v).programCounter = %d, should be %d", tt.programText, got.programCounter, tt.expectedProgramCounter)
		}
	}
}

var badOpcodePrograms = []struct {
	programText []int64
}{
	{[]int64{
		11, 0, 0, 0,
	}},
	{[]int64{
		-1, 0, 0, 0,
	}},
	{[]int64{
		98, 0, 0, 0,
	}},
}

func TestBadOpcode(t *testing.T) {
	for _, tt := range badOpcodePrograms {
		got := NewIntCodeProgram(tt.programText, nil, nil, nil)
		_, err := got.Execute()
		if err == nil {
			t.Errorf("NewIntCodeProgram(%v) opCode error not thrown", tt.programText)
		}
	}
}

var addOpcodePrograms = []struct {
	programText        []int64
	expectedFinalState []int64
}{
	{
		[]int64{
			1, 5, 6, 7,
			99, 25, 14, 0,
		},
		[]int64{
			1, 5, 6, 7,
			99, 25, 14, 39,
		},
	},
}

func TestAdditionOpcode(t *testing.T) {
	for _, tt := range addOpcodePrograms {
		got := NewIntCodeProgram(tt.programText, nil, nil, nil)
		_, err := got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !sliceutils.TestIntSliceEq(got.rawInstructions, tt.expectedFinalState) {
			t.Errorf("NewIntCodeProgram(%v) ended %v, not %v", tt.programText, got.rawInstructions, tt.expectedFinalState)
		}
	}
}

var multiplyOpcodePrograms = []struct {
	programText        []int64
	expectedFinalState []int64
}{
	{
		[]int64{
			2, 5, 6, 7,
			99, 25, 14, 0,
		},
		[]int64{
			2, 5, 6, 7,
			99, 25, 14, 350,
		},
	},
}

func TestMultiplyOpcode(t *testing.T) {
	for _, tt := range multiplyOpcodePrograms {
		got := NewIntCodeProgram(tt.programText, nil, nil, nil)
		_, err := got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !sliceutils.TestIntSliceEq(got.rawInstructions, tt.expectedFinalState) {
			t.Errorf("NewIntCodeProgram(%v) ended %v, not %v", tt.programText, got.rawInstructions, tt.expectedFinalState)
		}
	}
}

var echoProgram = []int64{3, 0, 4, 0, 99}

func TestInputAndOutputOperations(t *testing.T) {
	expectedValue := int64(67)

	input := make(chan int64)
	output := make(chan int64)
	final := make(chan int64)

	program := NewIntCodeProgram(echoProgram, input, output, final)

	finished := make(chan bool)

	go func(finished chan bool) {
		var outputValue = <-output
		if outputValue != expectedValue {
			t.Errorf("Output from echo program was %d, not %d", outputValue, expectedValue)
		}

		var finalValue = <-final
		if finalValue != expectedValue {
			t.Errorf("Final output from echo program was %d, not %d", finalValue, expectedValue)
		}
		finished <- true
	}(finished)

	go func(finished chan bool) {
		finalVal, err := program.Execute()
		if err != nil {
			t.Errorf("Error executing %v", err)
		}
		if finalVal != expectedValue {
			t.Errorf("Execute() got %d not %d", finalVal, expectedValue)
		}
	}(finished)

	go func(finished chan bool) {
		input <- expectedValue
	}(finished)

	<-finished
}

func TestFinalOperation(t *testing.T) {
	// input := make(chan int)
	// oupput := make(chan int)
	// final := make(chan int)
}

var complexPrograms = []struct {
	programText        []int64
	expectedFinalState []int64
}{
	{
		[]int64{
			1, 9, 10, 11,
			1, 11, 10, 9,
			99, 3, 4, 0,
		},
		[]int64{
			1, 9, 10, 11,
			1, 11, 10, 9,
			99, 11, 4, 7,
		},
	},
}

func TestComplexProgram(t *testing.T) {
	for _, tt := range complexPrograms {
		got := NewIntCodeProgram(tt.programText, nil, nil, nil)
		_, err := got.Execute()
		if err != nil {
			t.Errorf("NewIntCodeProgram(%v) error: %s", tt.programText, err)
		}

		if !sliceutils.TestIntSliceEq(got.rawInstructions, tt.expectedFinalState) {
			t.Errorf("NewIntCodeProgram(%v) ended %v, not %v", tt.programText, got.rawInstructions, tt.expectedFinalState)
		}
	}
}

var inputProgram = []int64{
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
	instruction        int64
	expectedOpcode     int64
	expectedModeParam1 int64
	expectedModeParam2 int64
	expectedModeParam3 int64
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
	instructionText []int64
	paramValue      int64
	paramMode       int64
	expectedValue   int64
}{
	{
		[]int64{0, 3, 2, 5},
		1,
		positional,
		3,
	},
	{
		[]int64{0, 3, 2, 5},
		1,
		immediate,
		1,
	},
	{
		[]int64{0, 3, 4, 5, 6},
		2,
		positional,
		4,
	},
	{
		[]int64{0, 3, 4, 5, 6},
		2,
		immediate,
		2,
	},
}

func TestFetchValue(t *testing.T) {
	for _, testData := range fetchValues {
		program := NewIntCodeProgram(testData.instructionText, nil, nil, nil)
		got, err := program.fetchValue(testData.paramValue, testData.paramMode)
		if err != nil {
			t.Errorf("Error thrown by FetchValue(%v) %v", testData.paramValue, err)
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

func TestFetchBadValue(t *testing.T) {
	program := NewIntCodeProgram([]int64{}, nil, nil, nil)
	_, err := program.fetchValue(int64(11), int64(3))
	if err == nil {
		t.Errorf("bad parameter mode ignored")
	}
}

func TestGetResult(t *testing.T) {
	var programIntructions = []int64{1, 2, 3, 4}
	program := NewIntCodeProgram(programIntructions, nil, nil, nil)
	got := program.GetResult()
	if got != 1 {
		t.Errorf("NewIntCodeProgram(%v).GetResult = %d, not 1", programIntructions, got)
	}
}

func TestSetInitialError(t *testing.T) {
	var programIntructions = []int64{1, 2, 3, 4}
	program := NewIntCodeProgram(programIntructions, nil, nil, nil)
	program.SetInitialError(123, 456)
	if program.rawInstructions[1] != 123 {
		t.Errorf("program.rawInstructions[1] = %d not 123", program.rawInstructions[1])
	}
	program.SetInitialError(123, 456)
	if program.rawInstructions[2] != 456 {
		t.Errorf("program.rawInstructions[2] = %d not 456", program.rawInstructions[1])
	}
}

func TestGetRelativeValue(t *testing.T) {
	copyProgram := []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	output := make(chan int64, len(copyProgram))
	program := NewIntCodeProgram(copyProgram, nil, output, nil)

	_, err := program.Execute()
	if err != nil {
		t.Errorf("NewIntCodeProgram(%v) error: %s", copyProgram, err)
	}

	var outputProg []int64
	for i := 0; i < len(copyProgram); i++ {
		outputProg = append(outputProg, <-output)
	}

	if len(outputProg) != len(copyProgram) {
		t.Errorf("Prog len differed %d != %d, %v", len(outputProg), len(copyProgram), outputProg)
	}

	for i := range outputProg {
		if outputProg[i] != copyProgram[i] {
			t.Errorf("OutputProg %v mistach at %d from copyProg %v", outputProg, i, copyProgram)
		}
	}

}

func TestLongNumbers(t *testing.T) {
	longNumProg := []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	expected := int64(1219070632396864)
	output := make(chan int64, 1)

	program := NewIntCodeProgram(longNumProg, nil, output, nil)

	result, err := program.Execute()
	if err != nil {
		t.Errorf("NewIntCodeProgram(%v) error: %s", longNumProg, err)
	}

	if result != expected {
		t.Errorf("NewIntCodeProgram(%v) = %d not %d", longNumProg, result, expected)
	}

	outputVal := <-output
	if outputVal != expected {
		t.Errorf("NewIntCodeProgram(%v) output = %d not %d", longNumProg, outputVal, expected)
	}

}

func TestExampleThree(t *testing.T) {
	longNumProg := []int64{104, 1125899906842624, 99}
	expected := int64(1125899906842624)
	output := make(chan int64, 1)

	program := NewIntCodeProgram(longNumProg, nil, output, nil)

	result, err := program.Execute()
	if err != nil {
		t.Errorf("NewIntCodeProgram(%v) error: %s", longNumProg, err)
	}

	if result != expected {
		t.Errorf("NewIntCodeProgram(%v) = %d not %d", longNumProg, result, expected)
	}

	outputVal := <-output
	if outputVal != expected {
		t.Errorf("NewIntCodeProgram(%v) output = %d not %d", longNumProg, outputVal, expected)
	}

}
