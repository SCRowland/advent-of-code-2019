package main

import (
	"challenge12"
	"fmt"
	"sort"
	"strings"
)

// LCM finds the lowest common multiple
func LCM(is []uint64) uint64 {
	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })

	biggestNum := is[len(is)-1]
	remaingingNums := is[:len(is)-2]

	fmt.Printf("%d %v\n", biggestNum, remaingingNums)

OUTER:
	for i := uint64(1); ; i++ {
		result := i * biggestNum

		for _, n := range remaingingNums {
			if result%n != 0 {
				continue OUTER
			}
		}

		return result
	}
}

func TimeToRepeat(input string) uint64 {
	moonsStrs := strings.Split(input, "\n")

	io := challenge12.NewMoon(moonsStrs[0])
	europa := challenge12.NewMoon(moonsStrs[1])
	ganymede := challenge12.NewMoon(moonsStrs[2])
	callisto := challenge12.NewMoon(moonsStrs[3])

	var initialPositions [4]challenge12.Position = [4]challenge12.Position{
		{io.Position.X, io.Position.Y, io.Position.Z},
		{europa.Position.X, europa.Position.Y, europa.Position.Z},
		{ganymede.Position.X, ganymede.Position.Y, ganymede.Position.Z},
		{callisto.Position.X, callisto.Position.Y, callisto.Position.Z},
	}

	var initialVelocities [4]challenge12.Velocity = [4]challenge12.Velocity{
		{io.Velocity.X, io.Velocity.Y, io.Velocity.Z},
		{europa.Velocity.X, europa.Velocity.Y, europa.Velocity.Z},
		{ganymede.Velocity.X, ganymede.Velocity.Y, ganymede.Velocity.Z},
		{callisto.Velocity.X, callisto.Velocity.Y, callisto.Velocity.Z},
	}

	var periods [4]uint64 = [4]uint64{}

	i := uint64(1)
	for ; ; i++ {
		io.ApplyGravityToBoth(europa)
		io.ApplyGravityToBoth(ganymede)
		io.ApplyGravityToBoth(callisto)

		europa.ApplyGravityToBoth(ganymede)
		europa.ApplyGravityToBoth(callisto)

		ganymede.ApplyGravityToBoth(callisto)

		io.ApplyVelocity()
		europa.ApplyVelocity()
		ganymede.ApplyVelocity()
		callisto.ApplyVelocity()

		if initialPositions[0] == io.Position && initialVelocities[0] == io.Velocity {
			if periods[0] == uint64(0) {
				periods[0] = i
			}
		}

		if initialPositions[1] == europa.Position && initialVelocities[1] == europa.Velocity {
			if periods[1] == uint64(0) {
				periods[1] = i
			}
		}

		if initialPositions[2] == ganymede.Position && initialVelocities[2] == ganymede.Velocity {
			if periods[2] == uint64(0) {
				periods[2] = i
			}
		}

		if initialPositions[3] == callisto.Position && initialVelocities[3] == callisto.Velocity {
			if periods[3] == uint64(0) {
				periods[3] = i
			}
		}

		if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 && periods[3] != 0 {
			break
		}
	}

	fmt.Printf("1=%d, 2=%d, 3=%d, 4=%d\n", periods[0], periods[1], periods[2], periods[3])
	answer := LCM(periods[:])

	return answer
}

func main() {
	output := TimeToRepeat(challenge12.Test1)
	fmt.Printf("OUTPUT1 %d\n", output)

	output = TimeToRepeat(challenge12.Test2)
	fmt.Printf("OUTPUT2 %d\n", output)

	output = TimeToRepeat(challenge12.PuzzleInput)
	fmt.Printf("OUTPUT3 %d\n", output)
}
