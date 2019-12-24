package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Running Intcode")

	convertedIntInstructions := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		rawInstructions := strings.Split(line, ",")

		for _, instruction := range rawInstructions {
			instructionCode, err := strconv.Atoi(instruction)
			if err != nil {
				fmt.Printf("err %s", err)
				continue
			}
			convertedIntInstructions = append(convertedIntInstructions, instructionCode)
		}

		for noun := 0; noun < 100; noun++ {
			for verb := 0; verb < 100; verb++ {
				rawProgram := make([]int, len(convertedIntInstructions))
				copy(rawProgram, convertedIntInstructions)
				program, err := NewIntCodeProgram(rawProgram)
				if err != nil {
					fmt.Printf("error creating program %s", err)
				}

				program.SetInitialError(noun, verb)
				program.Execute()
				finalState := program.rawInstructions[0]
				// fmt.Printf("Final State [%d] [%d] %v\n", i, j, finalState)
				if noun == 12 && verb == 2 {
					fmt.Printf("Final State noun=[%d] verb=[%d] answer=[%d] (%v)\n", noun, verb, (100*noun + verb), finalState)
				}

				if finalState == 19690720 {
					fmt.Printf("Found It:   noun=[%d] verb=[%d] answer=[%d] (%v)\n", noun, verb, (100*noun + verb), finalState)
					break
				}
			}
		}
	}
}
