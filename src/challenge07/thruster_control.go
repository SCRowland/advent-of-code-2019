package challenge07

import (
	"fmt"
	"intcode"
	"sliceutils"
	"strconv"
	"strings"
	"sync"
)

func connectChannels(input chan int, output chan int) {
	// pass everything from the input channel to the output channel
	for i := range input {
		output <- i
	}
}

func consumeChannel(channel chan int) {
	for range channel {
	}
}

func calculateThrust(programText string, phases [5]int, seedInput int, feedbackMode bool) (int, error) {
	// TODO confirm phases are all different
	var amplifiers [5]*intcode.Program

	var inputChannels [5]chan int
	var outputChannels [5]chan int
	var finalChannels [5]chan int
	for i := 0; i < 5; i++ {
		inputChannels[i] = make(chan int)
		outputChannels[i] = make(chan int)
		finalChannels[i] = make(chan int)
	}

	for i := 0; i < 5; i++ {
		convertedIntInstructions := []int{}
		rawInstructions := strings.Split(programText, ",")

		for _, instruction := range rawInstructions {
			instructionCode, err := strconv.Atoi(instruction)
			if err != nil {
				fmt.Printf("err %s", err)
			}
			convertedIntInstructions = append(convertedIntInstructions, instructionCode)
		}

		program := intcode.NewIntCodeProgram(convertedIntInstructions, inputChannels[i], outputChannels[i], finalChannels[i])
		amplifiers[i] = &program
	}

	// connect all channels
	for i := 0; i < 4; i++ {
		go connectChannels(outputChannels[i], inputChannels[i+1])
		go consumeChannel(finalChannels[i])
	}

	if feedbackMode {
		go connectChannels(outputChannels[4], inputChannels[0])
	} else {
		// just burn the final output channel!
		// go consumeChannel(outputChannels[4])
		go consumeChannel(finalChannels[4])
	}

	for i := 0; i < 5; i++ {
		go amplifiers[i].Execute()
	}

	var phaseWG sync.WaitGroup
	for i := 0; i < 5; i++ {
		phaseWG.Add(1)
		go func(channel chan int, phase int, wg *sync.WaitGroup) {
			channel <- phase
			wg.Done()
		}(inputChannels[i], phases[i], &phaseWG)
	}
	phaseWG.Wait()

	var seedWG sync.WaitGroup
	seedWG.Add(1)
	go func(channel chan int, seedInput int, wg *sync.WaitGroup) {
		channel <- seedInput
		wg.Done()
	}(inputChannels[0], seedInput, &seedWG)
	seedWG.Wait()

	var finalAnswers = <-outputChannels[4]

	return finalAnswers, nil
}

func calculateMaximumThrust(programText string, feedbackMode bool) (int, error) {
	maxThrust := 0

	var phasePossibilities [][]int

	if feedbackMode {
		phasePossibilities = sliceutils.Permutations([]int{5, 6, 7, 8, 9})
	} else {
		phasePossibilities = sliceutils.Permutations([]int{0, 1, 2, 3, 4})
	}

	for _, phasePossibility := range phasePossibilities {
		var phaseSettings [5]int
		var thrust int
		var err error
		copy(phaseSettings[:], phasePossibility[:5])

		thrust, err = calculateThrust(programText, [5]int(phaseSettings), 0, feedbackMode)
		if err != nil {
			return 0, fmt.Errorf("Error running: %v", err)
		}

		if thrust > maxThrust {
			maxThrust = thrust
		}
	}

	return maxThrust, nil
}
