package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Running Intcode")
	if len(os.Args) < 2 {
		fmt.Printf("\nPlease supply a program to run\n\n")
		return
	}
	fileName := os.Args[1]

	convertedIntInstructions := []int{}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("err %s", err)
	}
	rawInstructions := strings.Split(string(data), ",")

	for _, instruction := range rawInstructions {
		instructionCode, err := strconv.Atoi(instruction)
		if err != nil {
			fmt.Printf("err %s", err)
			return
		}
		convertedIntInstructions = append(convertedIntInstructions, instructionCode)
	}
	program := NewIntCodeProgram(convertedIntInstructions)
	err = program.Execute()
	if err != nil {
		fmt.Printf("There was an error: %v", err)
	}
	fmt.Printf("Program complete %d\n", program.GetResult())
}
