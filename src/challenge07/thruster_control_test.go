package challenge07

import (
	"testing"
)

func TestCalculateThrust(t *testing.T) {
	for _, testData := range examplePrograms {
		got, err := calculateThrust(testData.programText, testData.phaseSettings, 0, false)
		if err != nil {
			t.Errorf("calculateThrust(<>, %v) error %v", testData.phaseSettings, err)
		}
		if got != testData.maxthrust {
			t.Errorf("calculateThrust(<>, %v) = %d, not %d", testData.phaseSettings, got, testData.maxthrust)
		}
	}
}

func TestCalculateMaximumThrust(t *testing.T) {
	for _, testData := range examplePrograms {
		got, err := calculateMaximumThrust(testData.programText, false)
		if err != nil {
			t.Errorf("calculateMaximumThrust(<>) error %v", err)
		}
		if got != testData.maxthrust {
			t.Errorf("calculateMaximumThrust(<>) = %d, not %d", got, testData.maxthrust)
		}
	}
}

func TestSolveProblem(t *testing.T) {
	got, err := calculateMaximumThrust(challengeProgram, false)
	expected := int64(17440)
	if err != nil {
		t.Errorf("calculateMaximumThrust(<>) error %v", err)
	}
	if got != expected {
		t.Errorf("calculateMaximumThrust(<>) = %d, not %d", got, expected)
	}
}

func TestCalculateThrustFeedback(t *testing.T) {
	for _, testData := range exampleFeedbackPrograms {
		got, err := calculateThrust(testData.programText, testData.phaseSettings, 0, true)
		if err != nil {
			t.Errorf("calculateThrust(<>, %v) error %v", testData.phaseSettings, err)
		}
		if got != testData.maxthrust {
			t.Errorf("calculateThrust(<>, %v) = %d, not %d", testData.phaseSettings, got, testData.maxthrust)
		}
	}
}

func TestSolveProblemPartTwo(t *testing.T) {
	got, err := calculateMaximumThrust(challengeProgram, true)
	expected := int64(27561242)
	if err != nil {
		t.Errorf("calculateMaximumThrust(<>) error %v", err)
	}
	if got != expected {
		t.Errorf("calculateMaximumThrust(<>) = %d, not %d", got, expected)
	}
}
