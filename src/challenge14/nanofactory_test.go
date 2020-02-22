package challenge14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	testData := map[string]SampleInput{
		"example1":    example1,
		"example2":    example2,
		"example3":    example3,
		"example4":    example4,
		"example5":    example5,
		"puzzleInput": puzzleInput,
	}

	for tn, td := range testData {
		rm := parseReactionMap(td.input)
		actual := calcMinOre(rm, &Component{1, "FUEL"})

		assert.Equal(t, td.expected, actual, tn)
	}
}

func TestOneTrillionOre(t *testing.T) {
	testData := map[string]struct {
		input    string
		expected int64
	}{
		"example3": {
			example3.input,
			82892753,
		},
		"example4": {
			example4.input,
			5586022,
		},
		"example5": {
			example5.input,
			460664,
		},
		"puzzleInput": {
			puzzleInput.input,
			2074843,
		},
	}

	var oneTrillion int64 = 1000000000000

	for tn, td := range testData {
		rm := parseReactionMap(td.input)
		actual := maxFuel(rm, oneTrillion)

		assert.Equal(t, td.expected, actual, tn)
	}
}
