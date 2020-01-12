package challenge10

import "math"

// Angle is a float64 in radians
type Angle float64

// Distance is a gloat32
type Distance float32

// Point is a two dimensional point
type Point struct {
	x int
	y int
}

const angularDifferenceThreshold = Angle(0.000001)

// AngleEquals compares angles
func AngleEquals(a1, a2 Angle) bool {
	diff := a1 - a2
	diff = Angle(math.Abs(float64(diff)))

	if diff < angularDifferenceThreshold {
		return true
	}

	return false
}

// Angle calculates the angle between point p and another
func (p *Point) Angle(other *Point) Angle {
	adjLen := other.x - p.x
	oppLen := other.y - p.y

	radians := Angle(math.Atan2(float64(oppLen), float64(adjLen)))

	// (x > 0 ? x : (2*PI + x)) * 360 / (2*PI)

	var tempVal Angle
	if radians > 0 {
		tempVal = radians
	} else {
		tempVal = 2*math.Pi + radians
	}

	return tempVal * 360 / (2 * math.Pi)
}
