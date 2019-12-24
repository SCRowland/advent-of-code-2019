package main

import (
	"fmt"
)

const (
	add      = 1
	multiply = 2
	halt     = 99
)

type instructionRow struct {
	opcode        *int
	positionOne   *int
	positionTwo   *int
	positionThree *int
}

// IntCodeProgram represents the execution of an Int Code Program
type IntCodeProgram struct {
	programCounter  int
	rawInstructions []int
}

// NewIntCodeProgram is the preferred way to instantiate a New Int Code program
func NewIntCodeProgram(instructions []int) (IntCodeProgram, error) {
	program := IntCodeProgram{
		0, instructions,
	}
	return program, nil
}

// Execute executes the program
func (icp *IntCodeProgram) Execute() error {
	var opCodeInstructionCount = 0
	for ; icp.programCounter <= len(icp.rawInstructions); icp.programCounter += opCodeInstructionCount {
		var opcode = icp.rawInstructions[icp.programCounter]

		switch opcode {
		case halt:
			opCodeInstructionCount = 1
			return nil
		case add:
			opCodeInstructionCount = 4
			var positionOne = icp.rawInstructions[icp.programCounter+1]
			var positionTwo = icp.rawInstructions[icp.programCounter+2]
			var positionThree = icp.rawInstructions[icp.programCounter+3]
			icp.rawInstructions[positionThree] = icp.rawInstructions[positionOne] + icp.rawInstructions[positionTwo]
		case multiply:
			opCodeInstructionCount = 4
			var positionOne = icp.rawInstructions[icp.programCounter+1]
			var positionTwo = icp.rawInstructions[icp.programCounter+2]
			var positionThree = icp.rawInstructions[icp.programCounter+3]
			icp.rawInstructions[positionThree] = icp.rawInstructions[positionOne] * icp.rawInstructions[positionTwo]
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

func (icp *IntCodeProgram) SetInitialError(noun int, verb int) {
	icp.rawInstructions[1] = noun
	icp.rawInstructions[2] = verb
}
