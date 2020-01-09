package challenge07

import (
	"fmt"
	"intcode"
	"sliceutils"
	"strconv"
	"strings"
	"sync"
)

func connectChannels(input chan int64, output chan int64) {
	// pass everything from the input channel to the output channel
	for i := range input {
		output <- i
	}
}

func consumeChannel(channel chan int64) {
	for range channel {
	}
}

func calculateThrust(programText string, phases [5]int64, seedInput int64, feedbackMode bool) (int64, error) {
	// TODO confirm phases are all different
	var amplifiers [5]*intcode.Program

	var inputChannels [5]chan int64
	var outputChannels [5]chan int64
	var finalChannels [5]chan int64
	for i := 0; i < 5; i++ {
		inputChannels[i] = make(chan int64)
		outputChannels[i] = make(chan int64)
		finalChannels[i] = make(chan int64)
	}

	for i := 0; i < 5; i++ {
		convertedIntInstructions := []int64{}
		rawInstructions := strings.Split(programText, ",")

		for _, instruction := range rawInstructions {
			instructionCode, err := strconv.ParseInt(instruction, 10, 64)
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
		go func(channel chan int64, phase int64, wg *sync.WaitGroup) {
			channel <- phase
			wg.Done()
		}(inputChannels[i], phases[i], &phaseWG)
	}
	phaseWG.Wait()

	var seedWG sync.WaitGroup
	seedWG.Add(1)
	go func(channel chan int64, seedInput int64, wg *sync.WaitGroup) {
		channel <- seedInput
		wg.Done()
	}(inputChannels[0], seedInput, &seedWG)
	seedWG.Wait()

	if feedbackMode {
		var finalAnswer = int64(0)
		for i := range finalChannels[4] {
			finalAnswer = i
		}
		return finalAnswer, nil

	}

	var finalAnswer = <-outputChannels[4]
	return finalAnswer, nil
}

func calculateMaximumThrust(programText string, feedbackMode bool) (int64, error) {
	maxThrust := int64(0)

	var phasePossibilities [][]int64

	if feedbackMode {
		phasePossibilities = sliceutils.Permutations([]int64{5, 6, 7, 8, 9})
	} else {
		phasePossibilities = sliceutils.Permutations([]int64{0, 1, 2, 3, 4})
	}

	for _, phasePossibility := range phasePossibilities {
		var phaseSettings [5]int64
		var thrust int64
		var err error
		copy(phaseSettings[:], phasePossibility[:5])

		thrust, err = calculateThrust(programText, [5]int64(phaseSettings), 0, feedbackMode)
		if err != nil {
			return 0, fmt.Errorf("Error running: %v", err)
		}

		if thrust > maxThrust {
			maxThrust = thrust
		}
	}

	return maxThrust, nil
}
