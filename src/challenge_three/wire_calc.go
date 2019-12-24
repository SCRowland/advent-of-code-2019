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

// DistanceToPoint returns the distance the wire travels to a given point
func (w *Wire) DistanceToPoint(target Point) int {
	points, err := w.Points()
	if err != nil {
		// TODO error handling
	}
	for idx, p := range points {
		if target.x == p.x && target.y == p.y {
			return idx
		}
	}

	return 0
}

// NonOriginIntersections returns list of points where the wires cross
func NonOriginIntersections(wireOne, wireTwo Wire) []Point {
	points := []Point{}

	wireOnePoints, _ := wireOne.Points()
	wireTwoPoints, _ := wireTwo.Points()

	shortestList := []Point{}
	longestList := []Point{}

	if len(wireOnePoints) < len(wireTwoPoints) {
		shortestList = wireOnePoints
		longestList = wireTwoPoints
	} else {
		shortestList = wireTwoPoints
		longestList = wireOnePoints
	}

	shortestListPointsMap := make(map[Point]bool)
	for _, p1 := range shortestList {
		if p1.x == 0 && p1.y == 0 {
			// ignore origin
			continue
		}
		shortestListPointsMap[p1] = true
	}

	for _, p2 := range longestList {
		_, exists := shortestListPointsMap[p2]
		if exists {
			points = append(points, p2)
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
func NearestJunction(wireOneInstructions, wireTwoInstructions string) int {
	wireOne := NewWire(wireOneInstructions)
	wireTwo := NewWire(wireTwoInstructions)

	intersections := NonOriginIntersections(wireOne, wireTwo)
	nearestJunction, err := NearestPoint(intersections)
	if err != nil {
		// TODO implement!
	}

	return GetManhattenDistance(nearestJunction)
}

// StepsToJunction counts the steps to a junction
func StepsToJunction(wireOne, wireTwo Wire, junction Point) int {
	// TODO error checking!
	return wireOne.DistanceToPoint(junction) + wireTwo.DistanceToPoint(junction)
}

// LeastSteps returns the least number of steps to reach a junction
func LeastSteps(wireOneInstructions, wireTwoInstructions string) int {
	wireOne := NewWire(wireOneInstructions)
	wireTwo := NewWire(wireTwoInstructions)

	intersections := NonOriginIntersections(wireOne, wireTwo)

	// TODO are there more than one intersection?

	leastSteps := StepsToJunction(wireOne, wireTwo, intersections[0])
	for _, intersection := range intersections[1:] {
		steps := StepsToJunction(wireOne, wireTwo, intersection)
		if steps < leastSteps {
			leastSteps = steps
		}
	}

	return leastSteps
}
