package challenge12

import (
	"fmt"
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

	ganymede.applyGravityToBoth(callisto)

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

	europa.applyVelocity()

	assert.Equal(t, -1, europa.position.x)
	assert.Equal(t, 2, europa.position.y)
	assert.Equal(t, 6, europa.position.z)
}

func TestExample1(t *testing.T) {
	moonsStrs := strings.Split(test1, "\n")

	io := NewMoon(moonsStrs[0])
	europa := NewMoon(moonsStrs[1])
	ganymede := NewMoon(moonsStrs[2])
	callisto := NewMoon(moonsStrs[3])

	for i := 0; i < 10; i++ {
		io.applyGravityToBoth(europa)
		io.applyGravityToBoth(ganymede)
		io.applyGravityToBoth(callisto)

		europa.applyGravityToBoth(ganymede)
		europa.applyGravityToBoth(callisto)

		ganymede.applyGravityToBoth(callisto)

		io.applyVelocity()
		europa.applyVelocity()
		ganymede.applyVelocity()
		callisto.applyVelocity()
	}

	fmt.Println(io.toStr())
	fmt.Println(europa.toStr())
	fmt.Println(ganymede.toStr())
	fmt.Println(callisto.toStr())

	energy := CalculateTotalEnergy([]*Moon{io, europa, ganymede, callisto})
	assert.Equal(t, 179, energy)
}

func TestPuzzle(t *testing.T) {
	moonsStrs := strings.Split(puzzleInput, "\n")

	io := NewMoon(moonsStrs[0])
	europa := NewMoon(moonsStrs[1])
	ganymede := NewMoon(moonsStrs[2])
	callisto := NewMoon(moonsStrs[3])

	for i := 0; i < 1000; i++ {
		io.applyGravityToBoth(europa)
		io.applyGravityToBoth(ganymede)
		io.applyGravityToBoth(callisto)

		europa.applyGravityToBoth(ganymede)
		europa.applyGravityToBoth(callisto)

		ganymede.applyGravityToBoth(callisto)

		io.applyVelocity()
		europa.applyVelocity()
		ganymede.applyVelocity()
		callisto.applyVelocity()
	}

	fmt.Println(io.toStr())
	fmt.Println(europa.toStr())
	fmt.Println(ganymede.toStr())
	fmt.Println(callisto.toStr())

	energy := CalculateTotalEnergy([]*Moon{io, europa, ganymede, callisto})
	assert.Equal(t, 13500, energy)
}
