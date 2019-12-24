package main

import (
	"strings"
	"testing"
)

var testWireOneInstructions = "R8,U5,L5,D3"
var testWireTwoInstructions = "U7,R6,D4,L4"

var simpleTestInstructions = []struct {
	instruction string
	points      []Point
}{
	{
		"R2",
		[]Point{
			{0, 0},
			{1, 0},
			{2, 0},
		},
	},
	{
		"D2",
		[]Point{
			{0, 0},
			{0, -1},
			{0, -2},
		},
	},
	{
		"L2",
		[]Point{
			{0, 0},
			{-1, 0},
			{-2, 0},
		},
	},
	{
		"U2",
		[]Point{
			{0, 0},
			{0, 1},
			{0, 2},
		},
	},
}

func TestSimplePointLists(t *testing.T) {
	for _, tt := range simpleTestInstructions {
		wire := NewWire(tt.instruction)
		got, _ := wire.Points()

		if !EqualPoints(got, tt.points) {
			t.Errorf("NewWire(\"%s\").Points() = %v, not %v", tt.instruction, got, tt.points)
		}
	}
}

var unequalPointLists = []struct {
	listOne []Point
	listTwo []Point
}{
	{
		[]Point{
			{1, 3}, {4, 1},
		},
		[]Point{
			{1, 3},
		},
	},
	{
		[]Point{
			{1, 3}, {4, 1},
		},
		[]Point{
			{1, 3}, {4, 0},
		},
	},
}

func TestUnequalPointLists(t *testing.T) {
	for _, tt := range unequalPointLists {
		got := EqualPoints(tt.listOne, tt.listTwo)
		if got {
			t.Errorf("EqualPoints(%v, %v) = %t, should be false", tt.listOne, tt.listTwo, got)
		}
	}
}

var badInstructions = []struct {
	badInstructions string
	shouldContain   string
}{
	{
		"F45",
		"Unrecognised instruction",
	},
	{
		"f45",
		"Unrecognised instruction",
	},
	{
		"45",
		"Unrecognised instruction",
	},
	{
		"R2,U5,x",
		"Error processing instruction",
	},
	{
		"R",
		"Error processing instruction",
	},
	{
		"u",
		"Error processing instruction",
	},
}

func TestBadInstructions(t *testing.T) {
	for _, tt := range badInstructions {
		wire := NewWire(tt.badInstructions)
		_, err := wire.Points()

		if err == nil {
			t.Errorf("NewWire(\"%s\").Points() should error", tt.badInstructions)
		}

		if !strings.Contains(err.Error(), tt.shouldContain) {
			t.Errorf("NewWire(\"%s\").Points() error = %s, should contain %s", tt.shouldContain, err, tt.shouldContain)
		}
	}
}

var compoundTestInstructions = []struct {
	instruction string
	points      []Point
}{
	{
		"R2,L2",
		[]Point{
			{0, 0},
			{1, 0},
			{2, 0},
			{1, 0},
			{0, 0},
		},
	},
	{
		"U2,D2",
		[]Point{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 1},
			{0, 0},
		},
	},
}

func TestCompoundPointLists(t *testing.T) {
	for _, tt := range compoundTestInstructions {
		wire := NewWire(tt.instruction)
		got, _ := wire.Points()

		if !EqualPoints(got, tt.points) {
			t.Errorf("NewWire(\"%s\").Points() = %v, not %v", tt.instruction, got, tt.points)
		}
	}
}

func TestCachedPoints(t *testing.T) {
	wire := NewWire("")

	expectedPoints := []Point{
		{666, 777},
		{777, 666},
	}

	wire.points = expectedPoints

	got, _ := wire.Points()

	if !EqualPoints(got, wire.points) {
		t.Errorf("NewWire(<cached_value>).Points() = %v, not %v", got, expectedPoints)
	}
}

var nonIntersectingWires = []struct {
	wireOne       string
	wireTwo       string
	intersections []Point
}{
	{
		"R2,U2",
		"L2,D2",
		[]Point{},
	},
	{
		"U2,D2",
		"D2,R2",
		[]Point{},
	},
}

func TestNonInstersectingWires(t *testing.T) {
	for _, tt := range nonIntersectingWires {
		wireOne := NewWire(tt.wireOne)
		wireTwo := NewWire(tt.wireTwo)

		got := NonOriginIntersections(wireOne, wireTwo)

		if !EqualPoints(got, tt.intersections) {
			t.Errorf("NonOriginIntersections(\"%s\", \"%s\") = %v, not %v", tt.wireOne, tt.wireTwo, got, tt.intersections)
		}
	}
}

var intersectingWires = []struct {
	wireOne       string
	wireTwo       string
	intersections []Point
}{
	{
		"R2,U4",
		"U2,R4",
		[]Point{
			{2, 2},
		},
	},
	{
		"R2,U4,R4",
		"U2,R4,U4",
		[]Point{
			{2, 2},
			{4, 4},
		},
	},
}

func TestInstersectingWires(t *testing.T) {
	for _, tt := range intersectingWires {
		wireOne := NewWire(tt.wireOne)
		wireTwo := NewWire(tt.wireTwo)

		got := NonOriginIntersections(wireOne, wireTwo)

		if !EqualPoints(got, tt.intersections) {
			t.Errorf("NonOriginIntersections(\"%s\", \"%s\") = %v, not %v", tt.wireOne, tt.wireTwo, got, tt.intersections)
		}
	}
}

func TestEmptyListNearestIntersection(t *testing.T) {
	_, err := NearestPoint([]Point{})
	if err == nil {
		t.Errorf("NearestPoint([]Point{}) did not throw error")
	}

	if !strings.Contains(err.Error(), "No points supplied") {
		t.Errorf("NearestPoint([]Point{}) error = %s, should contain \"No points supplied\"", err)
	}
}

var pointList = []struct {
	point    Point
	distance int
}{
	{
		Point{1, 2},
		3,
	},
	{
		Point{-1, -2},
		3,
	},
	{
		Point{1, -2},
		3,
	},
	{
		Point{-1, 2},
		3,
	},
	{
		Point{1, 1},
		2,
	},
	{
		Point{-2013, 333},
		2346,
	},
}

func TestGetManhattenDistances(t *testing.T) {
	for _, tt := range pointList {
		got := GetManhattenDistance(tt.point)
		if got != tt.distance {
			t.Errorf("GetDistance(%v) = %d, not %d", tt.point, got, tt.distance)
		}
	}
}

var intersections = []struct {
	intersections []Point
	nearest       Point
}{
	{
		[]Point{
			{3, 3},
			{2, 2},
			{1, 1},
		},
		Point{1, 1},
	},
	{
		[]Point{
			{64, 34},
			{63, 34},
			{64, 35},
		},
		Point{63, 34},
	},
}

func TestNearestIntersection(t *testing.T) {
	for _, tt := range intersections {
		got, err := NearestPoint(tt.intersections)
		if err != nil {
			t.Errorf("NearestPoint(%v) err %s", tt.intersections, err.Error())
		}
		if got.x != tt.nearest.x && got.y != tt.nearest.y {
			t.Errorf("NearestPoint(%v) = %v, not %v", tt.intersections, got, tt.nearest)
		}
	}
}

func TestExampleOne(t *testing.T) {
	var wireOne = "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	var wireTwo = "U62,R66,U55,R34,D71,R55,D58,R83"
	var expectedDistance = 159

	var got = NearestJunction(wireOne, wireTwo)
	if got != expectedDistance {
		t.Errorf("NearestJunction(\"%s\", \"%s\") = %d, not %d", wireOne, wireTwo, got, expectedDistance)
	}
}

func TestExampleTwo(t *testing.T) {
	var wireOne = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	var wireTwo = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	var expectedDistance = 135

	var got = NearestJunction(wireOne, wireTwo)
	if got != expectedDistance {
		t.Errorf("NearestJunction(\"%s\", \"%s\") = %d, not %d", wireOne, wireTwo, got, expectedDistance)
	}
}
