package challenge11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaintHull(t *testing.T) {
	hull := NewHull()
	hull.PaintHull()

	t.Errorf("ASDSD")
}

func TestPaintedBounds(t *testing.T) {
	var testData = []struct {
		testPath []ScreenPosition
		minX     int64
		maxX     int64
		minY     int64
		maxY     int64
	}{
		{
			[]ScreenPosition{},
			int64(0),
			int64(0),
			int64(0),
			int64(0),
		},
		{
			[]ScreenPosition{
				{13, 0},
				{-13, 0},
				{0, 99},
				{0, -99},
			},
			int64(-13),
			int64(13),
			int64(-99),
			int64(99),
		},
	}

	for _, td := range testData {
		hull := NewHull()
		for _, sp := range td.testPath {
			hull.surface[sp] = white
		}

		minX, maxX, minY, maxY := hull.PaintedBounds()

		assert.Equal(t, td.minX, minX, "minX mismatch")
		assert.Equal(t, td.maxX, maxX, "maxX mismatch")
		assert.Equal(t, td.minY, minY, "minY mismatch")
		assert.Equal(t, td.maxY, maxY, "maxY mismatch")
	}
}
