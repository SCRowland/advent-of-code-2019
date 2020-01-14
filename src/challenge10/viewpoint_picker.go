package challenge10

import (
	"fmt"
	"sort"
	"strings"
)

// Asteroid is the representation of an asteroid on the map
type Asteroid Point

const (
	emptySpace = '.'
	asteroid   = '#'
	base       = 'X'
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
			if row[x] == asteroid || row[x] == base {
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

type ByDistance struct {
	points []Point
	from   Point
}

func (a ByDistance) Len() int { return len(a.points) }
func (a ByDistance) Less(i, j int) bool {
	return a.from.Distance(&a.points[i]) < a.from.Distance(&a.points[j])
}
func (a ByDistance) Swap(i, j int) { a.points[i], a.points[j] = a.points[j], a.points[i] }

func (aMap *AsteroidMap) getLinesOfAsteroids(from *Asteroid) *map[float32][]Point {
	linesOfAsteroids := make(map[float32][]Point)

	origin := Point(*from)

	for _, ast := range aMap.asteroids {
		if ast.x == origin.x && ast.y == origin.y {
			continue
		}
		pointB := Point(ast)
		rawAngle := float32(origin.Angle(&pointB))
		pList := linesOfAsteroids[rawAngle]
		pList = append(pList, pointB)
		linesOfAsteroids[rawAngle] = pList
	}

	// TODO sort the lists of asteroids
	for angle, pList := range linesOfAsteroids {
		pointsByDist := ByDistance{
			pList,
			origin,
		}
		sort.Sort(ByDistance(pointsByDist))
		linesOfAsteroids[angle] = pointsByDist.points
	}

	return &linesOfAsteroids
}

func (aMap *AsteroidMap) getAsteroidDestructionOrder(from *Asteroid) []Asteroid {
	asteroidsInDestructionOrder := []Asteroid{}

	linesOfAsteroids := aMap.getLinesOfAsteroids(from)

	var angles []float64
	for angle := range *linesOfAsteroids {
		angles = append(angles, float64(angle))
	}

	sort.Float64s(angles)

	// go round the angles until all asteroids are consumed
	nAsteroidsToFind := len(aMap.asteroids) - 1
	for len(asteroidsInDestructionOrder) < nAsteroidsToFind {
		for i := 0; i < len(angles); i++ {
			asteroids := (*linesOfAsteroids)[float32(angles[i])]
			if len(asteroids) > 0 {
				asteroidsInDestructionOrder = append(asteroidsInDestructionOrder, Asteroid(asteroids[0]))
				if len(asteroids) > 1 {
					(*linesOfAsteroids)[float32(angles[i])] = asteroids[1:]
				} else {
					(*linesOfAsteroids)[float32(angles[i])] = []Point{}
				}
			}
		}
	}

	return asteroidsInDestructionOrder
}

// GetVisibilityCount gets an array of ints showing visiblity counts from each location
func (aMap *AsteroidMap) GetVisibilityCount(from *Asteroid) int {
	var lineOfSightAngles = aMap.getLinesOfAsteroids(from)
	return len(*lineOfSightAngles)
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

// GetViewPoint finds the best view point for a base,
// and the number of asteroids you can see from it
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
