package intcode

import (
	"fmt"
)

// Program represents the execution of an Int Code Program
type Program struct {
	programCounter  int64
	relativeBase    int64
	rawInstructions []int64
	Input           chan int64
	Output          chan int64
	Final           chan int64
}

func parseInstruction(instruction int64) (opcode, paramMode1, paramMode2, paramMode3 int64) {
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

func (icp *Program) getRelativeValue(relLoc int64) int64 {
	location := icp.relativeBase + relLoc

	instructionLen := int64(len(icp.rawInstructions))

	if location > instructionLen {
		shortFall := location - instructionLen
		var i int64 = 0
		for ; i <= shortFall; i++ {
			icp.rawInstructions = append(icp.rawInstructions, 0)
		}
	}

	return icp.rawInstructions[location]
}

func (icp *Program) getAbsoluteValue(location int64) int64 {
	instructionLen := int64(len(icp.rawInstructions))

	if location > instructionLen {
		shortFall := location - instructionLen
		var i int64 = 0
		for ; i <= shortFall; i++ {
			icp.rawInstructions = append(icp.rawInstructions, 0)
		}
	}

	return icp.rawInstructions[location]
}

func (icp *Program) fetchValue(location, mode int64) (int64, error) {
	switch mode {
	case positional:
		return icp.getAbsoluteValue(location), nil
	case immediate:
		return location, nil
	case relative:
		return icp.getRelativeValue(location), nil
	default:
		return 0, fmt.Errorf("Unrecognised parameter mode %d", mode)
	}
}

func (icp *Program) setRelativeValue(relLoc, val int64) {
	location := icp.relativeBase + relLoc

	instructionLen := int64(len(icp.rawInstructions))

	if location >= instructionLen {
		shortFall := location - instructionLen + 1
		var i int64 = 0
		for ; i <= shortFall; i++ {
			icp.rawInstructions = append(icp.rawInstructions, 0)
		}
	}

	icp.rawInstructions[location] = val
}

func (icp *Program) setAbsoluteValue(location, val int64) {
	instructionLen := int64(len(icp.rawInstructions))

	if location >= instructionLen {
		shortFall := location - instructionLen + 1
		var i int64 = 0
		for ; i <= shortFall; i++ {
			icp.rawInstructions = append(icp.rawInstructions, 0)
		}
	}

	icp.rawInstructions[location] = val
}

func (icp *Program) setValue(location, val, mode int64) error {
	switch mode {
	case positional:
		icp.setAbsoluteValue(location, val)
		return nil
	case relative:
		icp.setRelativeValue(location, val)
		return nil
	default:
		return fmt.Errorf("Unrecognised parameter mode %d", mode)
	}
}

func (icp *Program) getLocationOrZero(location int64) int64 {
	if location < int64(len(icp.rawInstructions)) {
		return icp.rawInstructions[location]
	}
	return 0
}

// Execute executes the program
func (icp *Program) Execute() (int64, error) {
	var finalOutputVal int64

	if icp.Final != nil {
		defer close(icp.Final)
	}

	for icp.programCounter <= int64(len(icp.rawInstructions)) {
		var opcode, paramMode1, paramMode2, paramMode3 = parseInstruction(icp.rawInstructions[icp.programCounter])

		var positionOne = icp.getLocationOrZero(int64(icp.programCounter + 1))
		var positionTwo = icp.getLocationOrZero(int64(icp.programCounter + 2))
		var positionThree = icp.getLocationOrZero(int64(icp.programCounter + 3))

		switch opcode {
		case halt:
			if icp.Final != nil {
				icp.Final <- finalOutputVal
			}
			return finalOutputVal, nil

		case add:
			icp.programCounter += 4
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			var val2, _ = icp.fetchValue(positionTwo, paramMode2)
			icp.setValue(positionThree, val1+val2, paramMode3)

		case multiply:
			icp.programCounter += 4
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			var val2, _ = icp.fetchValue(positionTwo, paramMode2)
			icp.setValue(positionThree, val1*val2, paramMode3)

		case input:
			icp.programCounter += 2
			var inputNumber = <-icp.Input
			icp.setValue(positionOne, inputNumber, paramMode1)

		case output:
			icp.programCounter += 2
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			icp.Output <- val1
			finalOutputVal = val1

		case jumpIfTrue:
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			var val2, _ = icp.fetchValue(positionTwo, paramMode2)

			if val1 != 0 {
				icp.programCounter = val2
			} else {
				icp.programCounter += 3
			}

		case jumpIfFalse:
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			var val2, _ = icp.fetchValue(positionTwo, paramMode2)

			if val1 == 0 {
				icp.programCounter = val2
			} else {
				icp.programCounter += 3
			}

		case lessThan:
			icp.programCounter += 4
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			var val2, _ = icp.fetchValue(positionTwo, paramMode2)

			if val1 < val2 {
				icp.setValue(positionThree, 1, paramMode3)
			} else {
				icp.setValue(positionThree, 0, paramMode3)
			}

		case equals:
			icp.programCounter += 4
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			var val2, _ = icp.fetchValue(positionTwo, paramMode2)

			if val1 == val2 {
				icp.setValue(positionThree, 1, paramMode3)
			} else {
				icp.setValue(positionThree, 0, paramMode3)
			}

		case adjustRelBase:
			icp.programCounter += 2
			var val1, _ = icp.fetchValue(positionOne, paramMode1)
			icp.relativeBase += val1

		default:
			return output, fmt.Errorf("Bad instruction %d", opcode)
		}
	}
	return output, fmt.Errorf("No halt instruction")
}

// GetResult gets the program result
func (icp *Program) GetResult() int64 {
	return icp.rawInstructions[int64(0)]
}

// SetInitialError allows an initial error code to be set
func (icp *Program) SetInitialError(noun int64, verb int64) {
	icp.rawInstructions[1] = noun
	icp.rawInstructions[2] = verb
}

// NewIntCodeProgram is the only way to instantiate a New Int Code program
func NewIntCodeProgram(instructions []int64) *Program {
	input := make(chan int64, 1)
	output := make(chan int64, 1000)
	final := make(chan int64, 1)

	return &Program{
		0,
		0,
		instructions,
		input,
		output,
		final,
	}
}
