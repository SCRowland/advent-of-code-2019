package challenge11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHullPaintingRobot(t *testing.T) {
	emergHPR := NewEmergencyHullPaintingRobot()

	assert.Equal(t, emergHPR.position.across, int64(0), "bad across position")
	assert.Equal(t, emergHPR.position.down, int64(0), "bad down position")
}

func TestFirstMove(t *testing.T) {
	expectedColourToPaint := white

	emergHPR := NewEmergencyHullPaintingRobot()
	colourToPaint, _ := emergHPR.Move(black)

	assert.Equal(t, colourToPaint, expectedColourToPaint, "bad colour")
	assert.Equal(t, emergHPR.position.across, int64(-1), "bad across position")
}

func TestTurn(t *testing.T) {
	var testData = []struct {
		startDirection    int64
		directionToTurn   int64
		expectedDirection int64
	}{
		{
			up,
			left90,
			left,
		},
		{
			left,
			left90,
			down,
		},
		{
			down,
			left90,
			right,
		},
		{
			right,
			left90,
			up,
		},
		{
			up,
			right90,
			right,
		},
		{
			right,
			right90,
			down,
		},
		{
			down,
			right90,
			left,
		},
		{
			left,
			right90,
			up,
		},
	}

	for _, td := range testData {
		emergHPR := NewEmergencyHullPaintingRobot()
		emergHPR.direction = td.startDirection

		emergHPR.turn(td.directionToTurn)

		assert.Equal(t, td.expectedDirection, emergHPR.direction)
	}
}

func TestMove(t *testing.T) {
	var testData = []struct {
		direction     int64
		startPosition ScreenPosition
		endPosition   ScreenPosition
	}{
		{
			up,
			ScreenPosition{},
			ScreenPosition{int64(0), int64(-1)},
		},
		{
			left,
			ScreenPosition{},
			ScreenPosition{int64(-1), int64(0)},
		},
		{
			down,
			ScreenPosition{},
			ScreenPosition{int64(0), int64(1)},
		},
		{
			right,
			ScreenPosition{},
			ScreenPosition{int64(1), int64(0)},
		},
	}

	for _, td := range testData {
		emergHPR := NewEmergencyHullPaintingRobot()
		emergHPR.direction = td.direction
		emergHPR.position = td.startPosition

		emergHPR.move()

		assert.Equal(t, td.endPosition, emergHPR.position)
	}
}
