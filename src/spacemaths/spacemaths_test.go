package spacemaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	got := GCD(uint64(54), uint64(24))

	assert.Equal(t, uint64(6), got)
}

func TestLCM(t *testing.T) {
	ints := []uint64{
		uint64(4),
		uint64(6),
		uint64(24),
	}
	got := LCM(LCM(ints[0], ints[1]), ints[2])

	assert.Equal(t, uint64(24), got)
}

func TestLCM2(t *testing.T) {
	ints := []uint64{
		uint64(24),
		uint64(6),
		uint64(4),
	}
	got := LCM(LCM(ints[0], ints[1]), ints[2])

	assert.Equal(t, uint64(24), got)
}

func TestLCM3(t *testing.T) {
	ints := []uint64{
		uint64(924),
		uint64(2772),
		uint64(924),
		uint64(2772),
	}
	got := LCM(LCM(ints[0], ints[1]), ints[2])

	assert.Equal(t, uint64(2772), got)
}
