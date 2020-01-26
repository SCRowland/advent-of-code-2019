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
		x    int64
		y    int64
		z    int64
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

		assert.Equal(t, moon.Position.X, td.x, name)
		assert.Equal(t, moon.Position.Y, td.y, name)
		assert.Equal(t, moon.Position.Z, td.z, name)
	}
}

func TestToStr(t *testing.T) {
	testData := map[string]struct {
		expected string
		x        int64
		y        int64
		z        int64
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
			Position: Position{
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

	assert.Equal(t, int64(1), ganymede.Velocity.X)
	assert.Equal(t, int64(0), ganymede.Velocity.Y)
	assert.Equal(t, int64(-1), ganymede.Velocity.Z)

	assert.Equal(t, int64(-1), callisto.Velocity.X)
	assert.Equal(t, int64(0), callisto.Velocity.Y)
	assert.Equal(t, int64(1), callisto.Velocity.Z)
}

func TestApplyVelocity(t *testing.T) {
	europa := NewMoon("<x=1, y=2, z=3>")

	europa.Velocity.X = -2
	europa.Velocity.Y = 0
	europa.Velocity.Z = 3

	europa.ApplyVelocity()

	assert.Equal(t, int64(-1), europa.Position.X)
	assert.Equal(t, int64(2), europa.Position.Y)
	assert.Equal(t, int64(6), europa.Position.Z)
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
