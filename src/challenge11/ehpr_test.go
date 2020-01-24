package challenge11

import "testing"

func TestHullPaintingRobot(t *testing.T) {
	emergHPR := NewEmergencyHullPaintingRobot()

	emergHPR.program.Input <- 0
}
