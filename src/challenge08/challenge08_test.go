package challenge08

import "testing"

func TestNewImage(t *testing.T) {
	image := NewImage(sampleInput, sampleInputHeight, sampleInputWidth)
	for i := 0; i < len(image.layers); i++ {
		if image.layers[i] != sampleInputLayers[i] {
			t.Errorf("Layer %d is %s, should be %s", i, image.layers[i], sampleInputLayers[i])
		}
	}
}

func TestGetElfChecksum(t *testing.T) {
	image := NewImage(puzzleInput, 6, 25)
	expected := 1452

	if image.GetElfChecksum() != expected {
		t.Errorf("image.GetElfChecksum() = %d, not %d", image.GetElfChecksum(), expected)
	}

}
