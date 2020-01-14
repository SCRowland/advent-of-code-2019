package challenge10

import (
	"strings"
	"testing"
)

func PrintMap(aMap *AsteroidMap) {
	aMap.DebugPrint()
}

func TestExample1(t *testing.T) {
	aMap := NewAsteroidMap(example1Map, 5, 5)
	expectedAsteroids := 10

	if aMap.height != 5 || aMap.width != 5 {
		t.Errorf(
			"asteroid map has bad dimensions %dx%d not %dx%d",
			aMap.height, aMap.width, 5, 5,
		)
	}

	if len(aMap.asteroids) != expectedAsteroids {
		PrintMap(aMap)
		t.Errorf("asteroid map has wrong n asteroids got %d, not %d", len(aMap.asteroids), expectedAsteroids)
	}
}

func TestVisibilityCounts(t *testing.T) {
	aMap := NewAsteroidMap(example1Map, 5, 5)

	visCounts := aMap.GetVisibilityCounts()

	if len(visCounts) != aMap.height {
		PrintMap(aMap)
		t.Errorf("Bad visCounts len %d not %d", len(visCounts), aMap.height)
	}

	for y, row := range visCounts {
		for x := 0; x < len(row); x++ {
			if visCounts[y][x] != example1VisibilityCounts[y][x] {
				PrintMap(aMap)
				t.Errorf("Bad visCounts got %v not %v", visCounts, example1VisibilityCounts)
			}
		}
	}
}

func TestFindBestViewPoint(t *testing.T) {
	aMap := NewAsteroidMap(example1Map, 5, 5)

	bestViewPoint, visibleAsteroids := aMap.GetViewPoint()
	if bestViewPoint.x != example1X || bestViewPoint.y != example1Y {
		PrintMap(aMap)
		t.Errorf("Best view point is %v, not %d,%d", bestViewPoint, example1X, example1Y)
	}
	if visibleAsteroids != example1DetecableAsteroids {
		PrintMap(aMap)
		t.Errorf("Visible Asteroids = %d not %d", visibleAsteroids, example1DetecableAsteroids)
	}
}

func TestPrintf(t *testing.T) {
	// TODO - how to do this?
	aMap := NewAsteroidMap(example1Map, 5, 5)
	aMap.DebugPrint()
}

var allExamples = []struct {
	mapStr           string
	height           int
	width            int
	x                int
	y                int
	visibleAsteroids int
}{
	{
		example1Map,
		example1Height,
		example1Width,
		example1X,
		example1Y,
		example1DetecableAsteroids,
	},
	{
		example2Map,
		example2Height,
		example2Width,
		example2X,
		example2Y,
		example2DetecableAsteroids,
	},
	{
		example3Map,
		example3Height,
		example3Width,
		example3X,
		example3Y,
		example3DetecableAsteroids,
	},
	{
		example4Map,
		example4Height,
		example4Width,
		example4X,
		example4Y,
		example4DetecableAsteroids,
	},
	{
		example5Map,
		example5Height,
		example5Width,
		example5X,
		example5Y,
		example5DetecableAsteroids,
	},
}

func TestAllExamples(t *testing.T) {
	for _, example := range allExamples {
		aMap := NewAsteroidMap(example.mapStr, example.height, example.width)
		bestViewPoint, visibleAsteroids := aMap.GetViewPoint()
		if bestViewPoint.x != example.x || bestViewPoint.y != example.y {
			PrintMap(aMap)
			t.Errorf("Best view point is %v, not %d,%d", bestViewPoint, example.x, example.y)
		}
		if visibleAsteroids != example.visibleAsteroids {
			PrintMap(aMap)
			t.Errorf("Visible Asteroids = %d not %d", visibleAsteroids, example.visibleAsteroids)
		}
	}
}

func TestPuzzleInput(t *testing.T) {
	puzzleHeight := len(strings.Split(challenge10Input, "\n"))
	puzzleWidth := len(strings.Split(challenge10Input, "\n")[0])

	expectedX := 11
	expectedY := 13
	expectedVisibleAsteroids := 227

	aMap := NewAsteroidMap(challenge10Input, puzzleHeight, puzzleWidth)
	bestViewPoint, visibleAsteroids := aMap.GetViewPoint()
	if bestViewPoint.x != expectedX || bestViewPoint.y != expectedY {
		PrintMap(aMap)
		t.Errorf("Best view point is %v, not %d,%d", bestViewPoint, expectedX, expectedY)
	}
	if visibleAsteroids != expectedVisibleAsteroids {
		PrintMap(aMap)
		t.Errorf("Visible Asteroids = %d not %d", visibleAsteroids, expectedVisibleAsteroids)
	}
}

func TestPartTwoExample1(t *testing.T) {
	puzzleHeight := len(strings.Split(partTwoExample1.baseMap, "\n"))
	puzzleWidth := len(strings.Split(partTwoExample1.baseMap, "\n")[0])

	var expectedAsteroids = 37

	aMap := NewAsteroidMap(partTwoExample1.baseMap, puzzleHeight, puzzleWidth)
	if len(aMap.asteroids) != expectedAsteroids {
		PrintMap(aMap)
		t.Errorf("Map should have %d asteroids, not %d", expectedAsteroids, len(aMap.asteroids))
	}

	base := Asteroid{8, 3}
	orderOfDestruction := aMap.getAsteroidDestructionOrder(&base)

	first := Asteroid{8, 1}
	if orderOfDestruction[0] != first {
		PrintMap(aMap)
		t.Errorf("Bad first Asteroid %v, not %v", orderOfDestruction[0], first)
	}

	second := Asteroid{9, 0}
	if orderOfDestruction[1] != second {
		PrintMap(aMap)
		t.Errorf("Bad second Asteroid %v, not %v", orderOfDestruction[1], second)
	}
}

func TestPartTwoExample2(t *testing.T) {
	puzzleHeight := len(strings.Split(example5Map, "\n"))
	puzzleWidth := len(strings.Split(example5Map, "\n")[0])

	var expectedAsteroids = 300

	aMap := NewAsteroidMap(example5Map, puzzleHeight, puzzleWidth)
	if len(aMap.asteroids) != expectedAsteroids {
		PrintMap(aMap)
		t.Errorf("Map should have %d asteroids, not %d", expectedAsteroids, len(aMap.asteroids))
	}

	base := Asteroid{11, 13}
	orderOfDestruction := aMap.getAsteroidDestructionOrder(&base)

	first := Asteroid{11, 12}
	if orderOfDestruction[0] != first {
		PrintMap(aMap)
		t.Errorf("Bad first Asteroid %v, not %v", orderOfDestruction[0], first)
	}

	second := Asteroid{12, 1}
	if orderOfDestruction[1] != second {
		PrintMap(aMap)
		t.Errorf("Bad second Asteroid %v, not %v", orderOfDestruction[1], second)
	}

	if !(len(orderOfDestruction) > 200) {
		t.Errorf("Should be more than %d elements", len(orderOfDestruction))
	}

	twohundredth := Asteroid{8, 2}
	if orderOfDestruction[199] != twohundredth {
		t.Errorf("Bad 200th Asteroid %v, not %v", orderOfDestruction[199], twohundredth)
	}
}

func TestPartTwoMainPuzzle(t *testing.T) {
	puzzleHeight := len(strings.Split(challenge10Input, "\n"))
	puzzleWidth := len(strings.Split(challenge10Input, "\n")[0])

	var expectedAsteroids = 327

	aMap := NewAsteroidMap(challenge10Input, puzzleHeight, puzzleWidth)
	if len(aMap.asteroids) != expectedAsteroids {
		PrintMap(aMap)
		t.Errorf("Map should have %d asteroids, not %d", expectedAsteroids, len(aMap.asteroids))
	}

	base := Asteroid{11, 13}
	orderOfDestruction := aMap.getAsteroidDestructionOrder(&base)

	if !(len(orderOfDestruction) > 200) {
		t.Errorf("Should be more than %d elements", len(orderOfDestruction))
	}

	twohundredth := Asteroid{6, 4}
	if orderOfDestruction[199] != twohundredth {
		t.Errorf("Bad 200th Asteroid %v, not %v", orderOfDestruction[199], twohundredth)
	}
}
