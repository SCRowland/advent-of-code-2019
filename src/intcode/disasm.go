package intcode

// DisAsm disassembles the supplied byte array into strings
func (icp *Program) DisAsm() []string {
	var output []string

	for icp.programCounter <= int64(len(icp.rawInstructions)) {

		var opcode, _, _, _ = parseInstruction(icp.rawInstructions[icp.programCounter])

		var positions [3]int64
		positions[0] = icp.getLocationOrZero(int64(icp.programCounter + 1))
		positions[1] = icp.getLocationOrZero(int64(icp.programCounter + 2))
		positions[2] = icp.getLocationOrZero(int64(icp.programCounter + 3))

		// paramString := ""
		// for i := int64(0); i < (instructionLengths[opcode] - int64(1)); i++ {
		// 	paramString = fmt.Sprintf("%s %3d", paramString, positions[i])
		// }

		// output = append(output, fmt.Sprintf("%2d %s %s", len(output), instructionNames[opcode], paramString))
		// switch opcode {
		// case add:
		// 	var val1, _ = icp.fetchValue(positionOne, paramMode1)
		// 	var val2, _ = icp.fetchValue(positionTwo, paramMode2)
		// 	icp.setValue(positionThree, val1+val2, paramMode3)

		// case multiply:
		// 	var val1, _ = icp.fetchValue(positionOne, paramMode1)
		// 	var val2, _ = icp.fetchValue(positionTwo, paramMode2)
		// 	icp.setValue(positionThree, val1*val2, paramMode3)

		// case input:
		// 	var inputNumber = <-icp.input
		// 	icp.setValue(positionOne, inputNumber, paramMode1)

		// case output:
		// 	var val1, _ = icp.fetchValue(positionOne, paramMode1)
		// 	icp.output <- val1
		// 	finalOutputVal = val1

		// case jumpIfTrue:
		// 	var val1, _ = icp.fetchValue(positionOne, paramMode1)
		// 	var val2, _ = icp.fetchValue(positionTwo, paramMode2)

		// 	if val1 != 0 {
		// 		icp.programCounter = val2
		// 		continue
		// 	}

		// case jumpIfFalse:
		// 	var val1, _ = icp.fetchValue(positionOne, paramMode1)
		// 	var val2, _ = icp.fetchValue(positionTwo, paramMode2)

		// 	if val1 == 0 {
		// 		icp.programCounter = val2
		// 		continue
		// 	}

		// case lessThan:
		// 	var val1, _ = icp.fetchValue(positionOne, paramMode1)
		// 	var val2, _ = icp.fetchValue(positionTwo, paramMode2)

		// 	if val1 < val2 {
		// 		icp.setValue(positionThree, 1, paramMode3)
		// 	} else {
		// 		icp.setValue(positionThree, 0, paramMode3)
		// 	}

		// case equals:
		// 	var val1, _ = icp.fetchValue(positionOne, paramMode1)
		// 	var val2, _ = icp.fetchValue(positionTwo, paramMode2)

		// 	if val1 == val2 {
		// 		icp.setValue(positionThree, 1, paramMode3)
		// 	} else {
		// 		icp.setValue(positionThree, 0, paramMode3)
		// 	}

		// case adjustRelBase:
		// 	var val1, _ = icp.fetchValue(positionOne, paramMode1)
		// 	icp.relativeBase += val1

		// case halt:
		// 	if icp.final != nil {
		// 		icp.final <- finalOutputVal
		// 	}
		// 	return finalOutputVal, nil

		// default:
		// 	return output, fmt.Errorf("Bad instruction %d", opcode)
		// }
		if opcode == halt {
			return output
		}
		// icp.programCounter += instructionLengths[opcode]
	}

	return output
}
