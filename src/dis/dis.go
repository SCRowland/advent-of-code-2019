package main

import (
	"fmt"
	"intcode"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\nPlease supply a program to disassemble\n\n")
		return
	}

	fmt.Printf("disassembling intcode prog\n")

	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("err %s", err)
	}
	rawInstructions := strings.Split(string(data), ",")

	convertedIntInstructions := []int64{}
	for _, instruction := range rawInstructions {
		instructionCode, err := strconv.ParseInt(instruction, 10, 64)
		if err != nil {
			fmt.Printf("err %s", err)
		}
		convertedIntInstructions = append(convertedIntInstructions, instructionCode)
	}

	program := intcode.NewIntCodeProgram(convertedIntInstructions, nil, nil, nil)

	disassembly := program.DisAsm()
	for _, line := range disassembly {
		fmt.Println(line)
	}
}
