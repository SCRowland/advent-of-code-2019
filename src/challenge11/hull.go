package challenge11

import "fmt"

// Hull is a painting surface really
type Hull struct {
	surface map[ScreenPosition]colour
}

// PaintHull paints the hull with an emergency hull painting robot
func (h *Hull) PaintHull() {
	robot := NewEmergencyHullPaintingRobot()

	// move, draw, repeat
	// how to terminate? first loop 10541
	for {
		position := robot.position
		colourAtPosition := h.surface[position]

		newColour, stopped := robot.Move(colourAtPosition)
		if stopped {
			break
		}
		h.surface[position] = newColour
	}

	h.Draw()
	fmt.Printf("PANELS PAINTED %d\n", len(h.surface))
}

// PaintedBounds returns the bounds of the current painted hull
func (h *Hull) PaintedBounds() (minX, maxX, minY, maxY int64) {
	for sp := range h.surface {
		if sp.across < minX {
			minX = sp.across
		}
		if sp.across > maxX {
			maxX = sp.across
		}
		if sp.down < minY {
			minY = sp.down
		}
		if sp.down > maxY {
			maxY = sp.down
		}
	}

	return minX, maxX, minY, maxY
}

func translateToAxisPoint(p, minP, maxP int64) int64 {
	var offset int64
	if minP < 0 {
		offset = -minP
	} else {
		offset = minP
	}
	return p + offset
}

func translateFromAxisPoint(p, minP, maxP int64) int64 {
	return p + minP
}

// Draw draw the hull with current painting
func (h *Hull) Draw() {
	minX, maxX, minY, maxY := h.PaintedBounds()

	width := maxX - minX + 1
	height := maxY - minY + 1

	canvas := make([][]colour, width)
	for idx := range canvas {
		canvas[idx] = make([]colour, height)
	}

	for pt, col := range h.surface {
		across := translateToAxisPoint(pt.across, minX, maxX)
		down := translateToAxisPoint(pt.down, minY, maxY)
		canvas[across][down] = col
	}

	for x := int64(0); x < width; x++ {
		for y := int64(0); y < height; y++ {
			across := translateFromAxisPoint(x, minX, maxX)
			down := translateFromAxisPoint(y, minY, maxY)

			var toPrint rune
			if h.surface[ScreenPosition{across, down}] == white {
				toPrint = '#'
			} else {
				toPrint = ' '
			}
			fmt.Printf("%s", string(toPrint))
		}
		fmt.Println()
	}
	fmt.Println("================")
}

// NewHull instantiatiats a new hull
func NewHull() *Hull {
	paintedPath := make(map[ScreenPosition]colour)
	paintedPath[ScreenPosition{0, 0}] = black
	return &Hull{
		paintedPath,
	}
}
