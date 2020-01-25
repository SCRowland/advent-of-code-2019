package challenge12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
	z int
}

type Velocity struct {
	x int
	y int
	z int
}

// Moon is a moon with a three dimensional location
type Moon struct {
	position Position
	velocity Velocity
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
	moon.position.x = x
	moon.position.y = y
	moon.position.z = z

	return &moon
}

func (m *Moon) toStr() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>",
		m.position.x, m.position.y, m.position.z, m.velocity.x, m.velocity.y, m.velocity.z)
}

func (m *Moon) applyVelocity() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}

func (m *Moon) applyGravityToBoth(otherMoon *Moon) {
	if m.position.x == otherMoon.position.x {
	} else if m.position.x > otherMoon.position.x {
		m.velocity.x--
		otherMoon.velocity.x++
	} else {
		m.velocity.x++
		otherMoon.velocity.x--
	}

	if m.position.y == otherMoon.position.y {
	} else if m.position.y > otherMoon.position.y {
		m.velocity.y--
		otherMoon.velocity.y++
	} else {
		m.velocity.y++
		otherMoon.velocity.y--
	}

	if m.position.z == otherMoon.position.z {
	} else if m.position.z > otherMoon.position.z {
		m.velocity.z--
		otherMoon.velocity.z++
	} else {
		m.velocity.z++
		otherMoon.velocity.z--
	}
}

func NewMoon(repr string) *Moon {
	return fromStr(repr)
}
