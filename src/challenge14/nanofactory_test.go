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
