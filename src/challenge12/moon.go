package challenge12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
	Z int
}

type Velocity struct {
	X int
	Y int
	Z int
}

// Moon is a moon with a three dimensional location
type Moon struct {
	Position Position
	Velocity Velocity
}

func findCoord(coord, repr string) int {
	expr := fmt.Sprintf(`%s\s*=\s*-?\s*\d+`, coord)

	re := regexp.MustCompile(expr)
	match := re.Find([]byte(repr))

	re = regexp.MustCompile(`-?\s*\d+`)
	numberStr := re.Find(match)

	numberStr = []byte(strings.ReplaceAll(string(numberStr), " ", ""))

	number, _ := strconv.Atoi(string(numberStr))
	return number
}

func XYZfromStr(repr string) (x, y, z int) {
	x = findCoord("x", repr)
	y = findCoord("y", repr)
	z = findCoord("z", repr)

	return x, y, z
}

func fromStr(repr string) *Moon {
	moon := Moon{}

	x, y, z := XYZfromStr(repr)
	moon.Position.X = x
	moon.Position.Y = y
	moon.Position.Z = z

	return &moon
}

func (m *Moon) toStr() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>",
		m.Position.X, m.Position.Y, m.Position.Z, m.Velocity.X, m.Velocity.Y, m.Velocity.Z)
}

func (m *Moon) ShortStr() string {
	return strconv.Itoa(m.Position.X) + strconv.Itoa(m.Position.Y) + strconv.Itoa(m.Position.Z) + strconv.Itoa(m.Velocity.X) + strconv.Itoa(m.Velocity.Y) + strconv.Itoa(m.Velocity.Z)
}

func (m *Moon) ApplyVelocity() {
	m.Position.X += m.Velocity.X
	m.Position.Y += m.Velocity.Y
	m.Position.Z += m.Velocity.Z
}

func (m *Moon) ApplyGravityToBoth(otherMoon *Moon) {
	if m.Position.X == otherMoon.Position.X {
	} else if m.Position.X > otherMoon.Position.X {
		m.Velocity.X--
		otherMoon.Velocity.X++
	} else {
		m.Velocity.X++
		otherMoon.Velocity.X--
	}

	if m.Position.Y == otherMoon.Position.Y {
	} else if m.Position.Y > otherMoon.Position.Y {
		m.Velocity.Y--
		otherMoon.Velocity.Y++
	} else {
		m.Velocity.Y++
		otherMoon.Velocity.Y--
	}

	if m.Position.Z == otherMoon.Position.Z {
	} else if m.Position.Z > otherMoon.Position.Z {
		m.Velocity.Z--
		otherMoon.Velocity.Z++
	} else {
		m.Velocity.Z++
		otherMoon.Velocity.Z--
	}
}

func NewMoon(repr string) *Moon {
	return fromStr(repr)
}
