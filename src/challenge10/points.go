package challenge10

import (
	"math"
)

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

	var tempVal Angle
	if radians > 0 {
		tempVal = radians
	} else {
		tempVal = 2*math.Pi + radians
	}

	degreesFromHoriz := tempVal * 360.0 / (2 * math.Pi)

	degreesFromVerticalUp := degreesFromHoriz + 90
	if degreesFromVerticalUp >= 360 {
		return degreesFromVerticalUp - 360
	}
	return degreesFromVerticalUp
}

// Distance calculates the distance betweeo a point and the origin
func (p *Point) Distance(other *Point) float32 {
	// AB2 = dx2 + dy2
	dx := other.x - p.x
	dy := other.y - p.y

	distance := math.Sqrt(float64(dx*dx + dy*dy))

	return float32(distance)
}
