package challenge12

import (
	"fmt"
	"hash/fnv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReprInit(t *testing.T) {
	testData := map[string]struct {
		repr string
		x    int
		y    int
		z    int
	}{
		"Check read string": {
			"<x=1, y=2, z=3>",
			1,
			2,
			3,
		},
		"Check spaces ignored": {
			" < x = 1, y = 2, z = 3 > ",
			1,
			2,
			3,
		},
		"Check order ignored": {
			"<y=1, z=2, x=3>",
			3,
			1,
			2,
		},
	}

	for name, td := range testData {
		moon := NewMoon(td.repr)

		assert.Equal(t, moon.position.x, td.x, name)
		assert.Equal(t, moon.position.y, td.y, name)
		assert.Equal(t, moon.position.z, td.z, name)
	}
}

func TestToStr(t *testing.T) {
	testData := map[string]struct {
		expected string
		x        int
		y        int
		z        int
	}{
		"Test string repr": {
			"pos=<x=1, y=2, z=3>, vel=<x=0, y=0, z=0>",
			1,
			2,
			3,
		},
	}

	for name, td := range testData {
		moon := &Moon{
			position: Position{
				td.x,
				td.y,
				td.z,
			},
		}

		assert.Equal(t, td.expected, moon.toStr(), name)
	}
}

func TestApplyGravityTwoMoons(t *testing.T) {
	ganymede := NewMoon("<x=3, y=0, z=5>")
	callisto := NewMoon("<x=5, y=0, z=3>")

	ganymede.ApplyGravityToBoth(callisto)

	assert.Equal(t, 1, ganymede.velocity.x)
	assert.Equal(t, 0, ganymede.velocity.y)
	assert.Equal(t, -1, ganymede.velocity.z)

	assert.Equal(t, -1, callisto.velocity.x)
	assert.Equal(t, 0, callisto.velocity.y)
	assert.Equal(t, 1, callisto.velocity.z)
}

func TestApplyVelocity(t *testing.T) {
	europa := NewMoon("<x=1, y=2, z=3>")

	europa.velocity.x = -2
	europa.velocity.y = 0
	europa.velocity.z = 3

	europa.ApplyVelocity()

	assert.Equal(t, -1, europa.position.x)
	assert.Equal(t, 2, europa.position.y)
	assert.Equal(t, 6, europa.position.z)
}

func TestExample1(t *testing.T) {
	moonsStrs := strings.Split(Test1, "\n")

	io := NewMoon(moonsStrs[0])
	europa := NewMoon(moonsStrs[1])
	ganymede := NewMoon(moonsStrs[2])
	callisto := NewMoon(moonsStrs[3])

	for i := 0; i < 10; i++ {
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
	}

	fmt.Println(io.toStr())
	fmt.Println(europa.toStr())
	fmt.Println(ganymede.toStr())
	fmt.Println(callisto.toStr())

	energy := CalculateTotalEnergy([]*Moon{io, europa, ganymede, callisto})
	assert.Equal(t, 179, energy)
}

func TestPuzzlePartOne(t *testing.T) {
	moonsStrs := strings.Split(PuzzleInput, "\n")

	io := NewMoon(moonsStrs[0])
	europa := NewMoon(moonsStrs[1])
	ganymede := NewMoon(moonsStrs[2])
	callisto := NewMoon(moonsStrs[3])

	for i := 0; i < 1000; i++ {
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
	}

	fmt.Println(io.toStr())
	fmt.Println(europa.toStr())
	fmt.Println(ganymede.toStr())
	fmt.Println(callisto.toStr())

	energy := CalculateTotalEnergy([]*Moon{io, europa, ganymede, callisto})
	assert.Equal(t, 13500, energy)
}

var empty struct{}

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(temp1 int64, temp2 int64) {
	var lcmnum int64 = 1
	if temp1 > temp2 {
		lcmnum = temp1
	} else {
		lcmnum = temp2
	}
	/* Use of For Loop as a While Loop*/
	for {
		if lcmnum%temp1 == 0 && lcmnum%temp2 == 0 { // And operator
			/*  Print Statement with multiple variables   */
			fmt.Printf("LCM of %d and %d is %d", temp1, temp2, lcmnum)
			break
		}
		lcmnum++
	}
	return // Return without any value
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func TestExampleOne(t *testing.T) {
	moonsStrs := strings.Split(Test1, "\n")

	io := NewMoon(moonsStrs[0])
	europa := NewMoon(moonsStrs[1])
	ganymede := NewMoon(moonsStrs[2])
	callisto := NewMoon(moonsStrs[3])

	var initialPositions [4]Position = [4]Position{
		{io.position.x, io.position.y, io.position.z},
		{europa.position.x, europa.position.y, europa.position.z},
		{ganymede.position.x, ganymede.position.y, ganymede.position.z},
		{callisto.position.x, callisto.position.y, callisto.position.z},
	}

	var initialVelocities [4]Velocity = [4]Velocity{
		{io.velocity.x, io.velocity.y, io.velocity.z},
		{europa.velocity.x, europa.velocity.y, europa.velocity.z},
		{ganymede.velocity.x, ganymede.velocity.y, ganymede.velocity.z},
		{callisto.velocity.x, callisto.velocity.y, callisto.velocity.z},
	}

	var periods [4]int64 = [4]int64{}

	i := int64(1)
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

		if initialPositions[0] == io.position && initialVelocities[0] == io.velocity {
			if periods[0] == int64(0) {
				periods[0] = i
			}
		}

		if initialPositions[1] == europa.position && initialVelocities[1] == europa.velocity {
			if periods[1] == int64(0) {
				periods[1] = i
			}
		}

		if initialPositions[2] == ganymede.position && initialVelocities[2] == ganymede.velocity {
			if periods[2] == int64(0) {
				periods[2] = i
			}
		}

		if initialPositions[3] == callisto.position && initialVelocities[3] == callisto.velocity {
			if periods[3] == int64(0) {
				periods[3] = i
			}
		}

		if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 && periods[3] != 0 {
			break
		}
	}

	fmt.Printf("1: %d, 2: %d, 3: %d, 4: %d", periods[0], periods[1], periods[2], periods[3])
	answer := LCM(periods[0], periods[1], periods[2], periods[3])

	assert.Equal(t, int64(2772), answer)
}

func TestExample2(t *testing.T) {
	moonsStrs := strings.Split(Test2, "\n")

	io := NewMoon(moonsStrs[0])
	europa := NewMoon(moonsStrs[1])
	ganymede := NewMoon(moonsStrs[2])
	callisto := NewMoon(moonsStrs[3])

	var initialPositions [4]Position = [4]Position{
		{io.position.x, io.position.y, io.position.z},
		{europa.position.x, europa.position.y, europa.position.z},
		{ganymede.position.x, ganymede.position.y, ganymede.position.z},
		{callisto.position.x, callisto.position.y, callisto.position.z},
	}

	var initialVelocities [4]Velocity = [4]Velocity{
		{io.velocity.x, io.velocity.y, io.velocity.z},
		{europa.velocity.x, europa.velocity.y, europa.velocity.z},
		{ganymede.velocity.x, ganymede.velocity.y, ganymede.velocity.z},
		{callisto.velocity.x, callisto.velocity.y, callisto.velocity.z},
	}

	var periods [4]int64 = [4]int64{}

	i := int64(1)
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

		if initialPositions[0] == io.position && initialVelocities[0] == io.velocity {
			if periods[0] == int64(0) {
				periods[0] = i
			}
		}

		if initialPositions[1] == europa.position && initialVelocities[1] == europa.velocity {
			if periods[1] == int64(0) {
				periods[1] = i
			}
		}

		if initialPositions[2] == ganymede.position && initialVelocities[2] == ganymede.velocity {
			if periods[2] == int64(0) {
				periods[2] = i
			}
		}

		if initialPositions[3] == callisto.position && initialVelocities[3] == callisto.velocity {
			if periods[3] == int64(0) {
				periods[3] = i
			}
		}

		if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 && periods[3] != 0 {
			break
		}
	}

	fmt.Printf("1: %d, 2: %d, 3: %d, 4: %d", periods[0], periods[1], periods[2], periods[3])
	answer := LCM(periods[0], periods[1], periods[2], periods[3])

	assert.Equal(t, int64(2772), answer)
}

func TestPuzzleInput(t *testing.T) {
	moonsStrs := strings.Split(PuzzleInput, "\n")

	io := NewMoon(moonsStrs[0])
	europa := NewMoon(moonsStrs[1])
	ganymede := NewMoon(moonsStrs[2])
	callisto := NewMoon(moonsStrs[3])

	var initialPositions [4]Position = [4]Position{
		{io.position.x, io.position.y, io.position.z},
		{europa.position.x, europa.position.y, europa.position.z},
		{ganymede.position.x, ganymede.position.y, ganymede.position.z},
		{callisto.position.x, callisto.position.y, callisto.position.z},
	}

	var initialVelocities [4]Velocity = [4]Velocity{
		{io.velocity.x, io.velocity.y, io.velocity.z},
		{europa.velocity.x, europa.velocity.y, europa.velocity.z},
		{ganymede.velocity.x, ganymede.velocity.y, ganymede.velocity.z},
		{callisto.velocity.x, callisto.velocity.y, callisto.velocity.z},
	}

	var periods [4]int64 = [4]int64{}

	i := int64(1)
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

		if initialPositions[0] == io.position && initialVelocities[0] == io.velocity {
			if periods[0] == int64(0) {
				periods[0] = i
			}
		}

		if initialPositions[1] == europa.position && initialVelocities[1] == europa.velocity {
			if periods[1] == int64(0) {
				periods[1] = i
			}
		}

		if initialPositions[2] == ganymede.position && initialVelocities[2] == ganymede.velocity {
			if periods[2] == int64(0) {
				periods[2] = i
			}
		}

		if initialPositions[3] == callisto.position && initialVelocities[3] == callisto.velocity {
			if periods[3] == int64(0) {
				periods[3] = i
			}
		}

		if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 && periods[3] != 0 {
			break
		}
	}

	fmt.Printf("1: %d, 2: %d, 3: %d, 4: %d", periods[0], periods[1], periods[2], periods[3])
	answer := LCM(periods[0], periods[1], periods[2], periods[3])

	assert.Equal(t, int64(2772), answer)
}
