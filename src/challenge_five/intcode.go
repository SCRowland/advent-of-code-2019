package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	add      = 1
	multiply = 2
	input    = 3
	output   = 4
	halt     = 99
)

const (
	positional = 0
	immediate  = 1
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
	}
	return 0, nil
}

// Execute executes the program
func (icp *IntCodeProgram) Execute() error {
	var opCodeInstructionCount = 0
	for ; icp.programCounter <= len(icp.rawInstructions); icp.programCounter += opCodeInstructionCount {
		var instruction = icp.rawInstructions[icp.programCounter]

		opcode, paramMode1, paramMode2, _ := parseInstruction(instruction)

		switch opcode {
		case halt:
			opCodeInstructionCount = 1
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
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter integer: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("Bad input: %v", err)
			}
			inputNumber, err := strconv.Atoi(strings.Trim(text, "\n"))
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

func (icp *IntCodeProgram) SetInitialError(noun int, verb int) {
	icp.rawInstructions[1] = noun
	icp.rawInstructions[2] = verb
}
