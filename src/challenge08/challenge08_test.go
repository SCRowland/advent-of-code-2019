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

func TestSampleImage(t *testing.T) {
	sampleTwoInput := "0222112222120000"
	image := NewImage(sampleTwoInput, 2, 2)

	expected := "0110"

	expectedLayers := []string{
		"0222",
		"1122",
		"2212",
		"0000",
	}

	for i := 0; i < len(image.layers); i++ {
		if image.layers[i] != expectedLayers[i] {
			t.Errorf("Layer %d is %s, should be %s", i, image.layers[i], expectedLayers[i])
		}
	}

	render := image.GetRender()
	if render != expected {
		t.Errorf("Image doesn't match")
		image.Print()
	}
}

func TestGeneratedImage(t *testing.T) {
	image := NewImage(puzzleInput, 6, 25)

	expected := "111001001011100111101001010010100101001010000100101001011110100101110010010111001001011100100001001010000100101000010000100101000010010100001111001100"

	render := image.GetRender()
	if render != expected {
		t.Errorf("Image doesn't match")
		image.Print()
	}
}
