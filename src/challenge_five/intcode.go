package main

import (
	"fmt"
)

// IntCodeProgram represents the execution of an Int Code Program
type IntCodeProgram struct {
	programCounter  int
	rawInstructions []int
}

// NewIntCodeProgram is the only way to instantiate a New Int Code program
func NewIntCodeProgram(instructions []int) (IntCodeProgram, error) {
	program := IntCodeProgram{
		0, instructions,
	}
	return program, nil
}

func parseInstruction(instruction int) (opcode, paramMode1, paramMode2, paramMode3 int) {
	/*
			ABCDE
			1002

			DE - two-digit opcode,      02 == opcode 2
			 C - mode of 1st parameter,  0 == position mode
			 B - mode of 2nd parameter,  1 == immediate mode
			 A - mode of 3rd parameter,  0 == position mode,
		                                  omitted due to being a leading zero
	*/
	opcode = instruction % 100
	paramMode1 = instruction / 100 % 10
	paramMode2 = instruction / 1000 % 10
	paramMode3 = instruction / 10000 % 10

	return opcode, paramMode1, paramMode2, paramMode3
}

func (icp *IntCodeProgram) fetchValue(location, mode int) (int, error) {
	switch mode {
	case positional:
		return icp.rawInstructions[location], nil
	case immediate:
		return location, nil
	default:
		return 0, fmt.Errorf("Unrecognised parameter mode %d", mode)
	}
}

// Execute executes the program
func (icp *IntCodeProgram) Execute() error {
	for opCodeInstructionCount := 0; icp.programCounter <= len(icp.rawInstructions); icp.programCounter += opCodeInstructionCount {
		opcode, paramMode1, paramMode2, _ := parseInstruction(icp.rawInstructions[icp.programCounter])

		switch opcode {
		case halt:
			return nil

		case add:
			opCodeInstructionCount = 4
			var positionOne = icp.rawInstructions[icp.programCounter+1]
			var positionTwo = icp.rawInstructions[icp.programCounter+2]
			var positionThree = icp.rawInstructions[icp.programCounter+3]
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			var val2, _ = icp.fetchValue(positionTwo, paramMode2)
			icp.rawInstructions[positionThree] = val1 + val2

		case multiply:
			opCodeInstructionCount = 4
			var positionOne = icp.rawInstructions[icp.programCounter+1]
			var positionTwo = icp.rawInstructions[icp.programCounter+2]
			var positionThree = icp.rawInstructions[icp.programCounter+3]
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			var val2, _ = icp.fetchValue(positionTwo, paramMode2)
			icp.rawInstructions[positionThree] = val1 * val2

		case input:
			opCodeInstructionCount = 2
			var positionOne = icp.rawInstructions[icp.programCounter+1]
			fmt.Print("Enter integer: ")
			var inputNumber int
			_, err := fmt.Scanf("%d", &inputNumber)
			if err != nil {
				return fmt.Errorf("Bad input: %v", err)
			}
			icp.rawInstructions[positionOne] = inputNumber

		case output:
			opCodeInstructionCount = 2
			var positionOne = icp.rawInstructions[icp.programCounter+1]
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			fmt.Printf("OUTPUT: %d\n", val1)

		default:
			return fmt.Errorf("Bad instruction %d", opcode)
		}
	}
	return nil
}

// GetResult gets the program result
func (icp *IntCodeProgram) GetResult() int {
	return icp.rawInstructions[0]
}

// SetInitialError allows an initial error code to be set
func (icp *IntCodeProgram) SetInitialError(noun int, verb int) {
	icp.rawInstructions[1] = noun
	icp.rawInstructions[2] = verb
}
