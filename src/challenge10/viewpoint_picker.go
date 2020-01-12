package challenge10

import (
	"fmt"
	"strings"
)

// Asteroid is the representation of an asteroid on the map
type Asteroid Point

const (
	emptySpace = '.'
	asteroid   = '#'
)

// AsteroidMap keeps a map of the asteroids
type AsteroidMap struct {
	height int
	width  int

	rawMapRows []string
	asteroids  []Asteroid
}

// NewAsteroidMap creates a new asteroid map
func NewAsteroidMap(asteroidMapStr string, height, width int) *AsteroidMap {
	aMap := AsteroidMap{height, width, []string{}, []Asteroid{}}

	aMap.rawMapRows = strings.Split(asteroidMapStr, "\n")

	for y, row := range aMap.rawMapRows {
		for x := 0; x < width; x++ {
			if row[x] == asteroid {
				aMap.asteroids = append(aMap.asteroids, Asteroid{x, y})
			}
		}
	}

	return &aMap
}

// DebugPrint prints an Asteroid Map and some debug info
func (aMap *AsteroidMap) DebugPrint() {
	for _, row := range aMap.rawMapRows {
		fmt.Printf("[%s]\n", row)
	}

	var representation [][]bool

	for y := 0; y < aMap.height; y++ {
		var row []bool
		for x := 0; x < aMap.width; x++ {
			row = append(row, false)
		}
		representation = append(representation, row)
	}

	for _, a := range aMap.asteroids {
		representation[a.y][a.x] = true
	}

	for _, row := range representation {
		for i := 0; i < len(row); i++ {
			if row[i] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	visCounts := aMap.GetVisibilityCounts()
	for _, row := range visCounts {
		fmt.Printf("%v\n", row)
	}
}

// GetVisibilityCount gets an array of ints showing visiblity counts from each location
func (aMap *AsteroidMap) GetVisibilityCount(a *Asteroid) int {
	var lineOfSightAngles map[float32]bool = make(map[float32]bool)

	origin := Point(*a)

	for _, ast := range aMap.asteroids {
		if ast.x == a.x && ast.y == a.y {
			continue
		}
		pointB := Point(ast)
		rawAngle := origin.Angle(&pointB)
		lineOfSightAngles[float32(rawAngle)] = true
	}
	return len(lineOfSightAngles)
}

// GetVisibilityCounts gets an array of ints showing visiblity counts from each location
func (aMap *AsteroidMap) GetVisibilityCounts() [][]int {
	var visCounts [][]int

	for y := 0; y < aMap.height; y++ {
		var row []int
		for x := 0; x < aMap.width; x++ {
			row = append(row, 0)
		}
		visCounts = append(visCounts, row)
	}

	for _, a := range aMap.asteroids {
		visCounts[a.y][a.x] = aMap.GetVisibilityCount(&a)
	}

	return visCounts
}

// GetViewPoint finds the best view point for a base
func (aMap *AsteroidMap) GetViewPoint() (Asteroid, int) {
	visCounts := aMap.GetVisibilityCounts()

	maxViewCount := 0
	maxViewAsteroid := Asteroid{}

	for y := 0; y < aMap.height; y++ {
		row := visCounts[y]
		for x := 0; x < len(row); x++ {
			viewCount := visCounts[y][x]
			if viewCount > maxViewCount {
				maxViewCount = viewCount
				maxViewAsteroid = Asteroid{x, y}
			}
		}
	}

	return maxViewAsteroid, maxViewCount
}
