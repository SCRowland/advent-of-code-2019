package main

import (
	"challenge12"
	"fmt"
	"spacemaths"
	"strings"
)

func timeToRepeat(input string) uint64 {
	moonStrs := strings.Split(input, "\n")

	moons := [4]*challenge12.Moon{
		challenge12.NewMoon(moonStrs[0]),
		challenge12.NewMoon(moonStrs[1]),
		challenge12.NewMoon(moonStrs[2]),
		challenge12.NewMoon(moonStrs[3]),
	}

	periods := [3]uint64{
		uint64(0),
		uint64(0),
		uint64(0),
	}

	i := uint64(1)
	for ; ; i++ {
		// Gravity
		moons[0].ApplyGravityToBoth(moons[1])
		moons[0].ApplyGravityToBoth(moons[2])
		moons[0].ApplyGravityToBoth(moons[3])

		moons[1].ApplyGravityToBoth(moons[2])
		moons[1].ApplyGravityToBoth(moons[3])

		moons[2].ApplyGravityToBoth(moons[3])

		for _, moon := range moons {
			moon.ApplyVelocity()
		}

		if periods[0] == 0 {
			foundMovingX := false
			for _, moon := range moons {
				if moon.Velocity.X != 0 {
					foundMovingX = true
					break
				}
			}

			if !foundMovingX {
				periods[0] = i * 2
			}
		}

		if periods[1] == 0 {
			foundMovingY := false
			for _, moon := range moons {
				if moon.Velocity.Y != 0 {
					foundMovingY = true
					break
				}
			}

			if !foundMovingY {
				periods[1] = i * 2
			}
		}

		if periods[2] == 0 {
			foundMovingZ := false
			for _, moon := range moons {
				if moon.Velocity.Z != 0 {
					foundMovingZ = true
					break
				}
			}

			if !foundMovingZ {
				periods[2] = i * 2
			}
		}

		if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 {
			break
		}
	}

	fmt.Printf("1=%d, 2=%d, 3=%dn", periods[0], periods[1], periods[2])
	answer := spacemaths.LCM(spacemaths.LCM(periods[0], periods[1]), periods[2])

	return answer
}

func main() {
	output := timeToRepeat(challenge12.Test1)
	fmt.Printf("OUTPUT1 %d\n", output)

	output = timeToRepeat(challenge12.Test2)
	fmt.Printf("OUTPUT2 %d\n", output)

	output = timeToRepeat(challenge12.PuzzleInput)
	fmt.Printf("OUTPUT3 %d\n", output)
}
