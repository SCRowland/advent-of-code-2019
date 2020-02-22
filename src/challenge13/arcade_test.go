package challenge13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArcade(t *testing.T) {
	arcade := NewArcade()

	assert.NotNil(t, arcade)
}

func TestMove(t *testing.T) {
	arcade := NewArcade()

	across, down, thing := arcade.GetOutput()

	assert.Equal(t, int64(0), across)
	assert.Equal(t, int64(0), down)
	assert.Equal(t, int64(1), thing)
}

func TestBlockCount(t *testing.T) {
	arcade := NewArcade()

	blockCount := int64(0)

RUN:
	for {
		select {
		case <-arcade.program.Final:
			break RUN
		default:
		}
		_, _, thing := arcade.GetOutput()
		if thing == 2 {
			blockCount++
		}
	}
	assert.Equal(t, int64(462), blockCount)
}

func TestGameBounds(t *testing.T) {
	arcade := NewArcade()

	maxAcross := int64(0)
	maxDown := int64(0)

RUN:
	for {
		select {
		case <-arcade.program.Final:
			break RUN
		default:
		}
		across, down, _ := arcade.GetOutput()
		if across > maxAcross {
			maxAcross = across
		}
		if down > maxDown {
			maxDown = down
		}
	}
	assert.Equal(t, int64(44), maxAcross)
	assert.Equal(t, int64(23), maxDown)
}
