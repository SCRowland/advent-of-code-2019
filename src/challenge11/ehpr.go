// Package challenge11 is my solution to https://adventofcode.com/2019/day/11
package challenge11

import (
	"intcode"
)

// EmergencyHullPaintingRobot is an emergency hull painting robot
type EmergencyHullPaintingRobot struct {
	program *intcode.Program
}

// NewEmergencyHullPaintingRobot instantiates a EmergencyHullPaintingRobot
func NewEmergencyHullPaintingRobot() *EmergencyHullPaintingRobot {
	return &EmergencyHullPaintingRobot{
		intcode.NewIntCodeProgram(puzzleInput),
	}
}
