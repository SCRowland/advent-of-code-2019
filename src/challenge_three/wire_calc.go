package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Point represents a two dimensional point
type Point struct {
	x int
	y int
}

// Wire describes a wire track
type Wire struct {
	instructions string
	points       []Point
	length       int
}

// NewWire creates a new wire with supplied instructions
func NewWire(wireInstructions string) Wire {
	newWire := Wire{
		wireInstructions, nil, 0,
	}
	return newWire
}

// EqualPoints tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func EqualPoints(a, b []Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Points returns all points on a line
func (w *Wire) Points() ([]Point, error) {
	if w.points != nil {
		return w.points, nil
	}
	var points = []Point{
		Point{0, 0},
	}
	var x = 0
	var y = 0
	for _, instr := range strings.Split(w.instructions, ",") {
		direction := string([]rune(instr)[0])
		distance, err := strconv.Atoi(string([]rune(instr)[1:]))
		if err != nil {
			return nil, fmt.Errorf("Error processing instruction %s: %s", instr, err)
		}

		for i := 0; i < distance; i++ {
			switch direction {
			case "R":
				x++
			case "D":
				y--
			case "L":
				x--
			case "U":
				y++
			default:
				return nil, fmt.Errorf("Unrecognised instruction  %s", instr)
			}
			point := Point{x, y}
			points = append(points, point)
		}
	}
	w.points = points
	w.length = len(points)
	return points, nil
}

// NonOriginIntersections returns list of points where the wires cross
func NonOriginIntersections(l Wire, r Wire) []Point {
	points := []Point{}

	lPoints, _ := l.Points()
	rPoints, _ := r.Points()

	for _, l := range lPoints {
		for _, r := range rPoints {
			if l.x == 0 && l.y == 0 {
				// intersecting at the origin doesn't count!
				continue
			}
			if l.x == r.x && l.y == r.y {
				points = append(points, Point{l.x, l.y})
			}
		}
	}

	return points
}

// GetManhattenDistance returns absolute distance from origin to point
func GetManhattenDistance(point Point) int {
	x := point.x
	y := point.y

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	return x + y
}

// NearestPoint calculates which point is nearest the origin
func NearestPoint(points []Point) (Point, error) {
	if len(points) == 0 {
		return Point{}, errors.New("No points supplied")
	}

	nearestPoint := points[0]
	shortestDistance := GetManhattenDistance(points[0])

	for _, point := range points[1:] {
		distance := GetManhattenDistance(point)
		if distance < shortestDistance {
			shortestDistance = distance
			nearestPoint = point
		}
	}

	return nearestPoint, nil
}

// NearestJunction calculates the distance to the nearest junction to the origin
func NearestJunction(l string, r string) int {
	wireOne := NewWire(l)
	wireTwo := NewWire(r)

	intersections := NonOriginIntersections(wireOne, wireTwo)
	nearestJunction, err := NearestPoint(intersections)
	if err != nil {
		// TODO implement!
	}

	return GetManhattenDistance(nearestJunction)
}
