// Package challenge11 is my solution to https://adventofcode.com/2019/day/11
package challenge11

import (
	"fmt"
	"intcode"
)

// ScreenPosition is a position where 0,0 is the top left of the screen
type ScreenPosition struct {
	across int64
	down   int64
}

// EmergencyHullPaintingRobot is an emergency hull painting robot
type EmergencyHullPaintingRobot struct {
	program   *intcode.Program
	position  ScreenPosition
	direction int64
}

// NewEmergencyHullPaintingRobot instantiates a EmergencyHullPaintingRobot
func NewEmergencyHullPaintingRobot() *EmergencyHullPaintingRobot {
	program := intcode.NewIntCodeProgram(puzzleInput)

	go program.Execute()

	return &EmergencyHullPaintingRobot{
		program,
		ScreenPosition{},
		up,
	}
}

func (ehpr *EmergencyHullPaintingRobot) turn(directionToTurn int64) {
	switch directionToTurn {
	case left90:
		ehpr.direction--
		if ehpr.direction < 0 {
			ehpr.direction = left
		}
	case right90:
		ehpr.direction++
		if ehpr.direction > 3 {
			ehpr.direction = up
		}
	default:
		fmt.Printf("Bad direction %v", directionToTurn)
	}
}

func (ehpr *EmergencyHullPaintingRobot) move() {
	switch ehpr.direction {
	case up:
		ehpr.position.down--
	case right:
		ehpr.position.across++
	case down:
		ehpr.position.down++
	case left:
		ehpr.position.across--
	default:
		fmt.Printf("Bad direction %v", ehpr.direction)
	}
}

// Move moves the robot a single time
func (ehpr *EmergencyHullPaintingRobot) Move(currentColour colour) (colourToPaint colour, stopped bool) {
	ehpr.program.Input <- int64(currentColour)

	// stop if it's done
	select {
	case <-ehpr.program.Final:
		return colour(-1), true
	default:
	}

	colourToPaint = colour(<-ehpr.program.Output)
	directionToMove := <-ehpr.program.Output

	ehpr.turn(directionToMove)
	ehpr.move()

	return colour(colourToPaint), false
}
