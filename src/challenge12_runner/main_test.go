package main

import (
	"challenge12"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeToRepeat1(t *testing.T) {
	output := timeToRepeat(challenge12.Test1)
	assert.Equal(t, uint64(2772), output)
}

func TestTimeToRepeat2(t *testing.T) {
	output := timeToRepeat(challenge12.Test2)
	assert.Equal(t, uint64(4686774924), output)
}

func TestTimeToRepeatPuzzleInput(t *testing.T) {
	output := timeToRepeat(challenge12.PuzzleInput)
	assert.Equal(t, uint64(278013787106916), output)
}
