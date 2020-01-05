package challenge08

import (
	"regexp"
)

// Image represents an Image
type Image struct {
	data   string
	height int
	width  int
	layers []string
}

// NewImage creates an image
func NewImage(data string, height, width int) *Image {
	var layers []string
	layerLength := height * width

	var layerStart = 0
	for h := 0; h < height; h++ {
		layers = append(layers, data[layerStart:layerStart+layerLength])
		layerStart += layerLength
	}

	var image = Image{
		data,
		height,
		width,
		layers,
	}

	return &image
}

// GetElfChecksum sfsdf
func (img *Image) GetElfChecksum() int {
	var zeroCount []int

	nZeros := regexp.MustCompile("0")

	for i := 0; i < img.height; i++ {
		matches := nZeros.FindAllStringIndex(img.layers[i], -1)
		zeroCount = append(zeroCount, len(matches))
	}

	layer := 0
	minZeros := zeroCount[0]

	for i := 0; i < len(zeroCount); i++ {
		if zeroCount[i] < minZeros {
			minZeros = zeroCount[i]
			layer = i
		}
	}

	minZeroLayer := img.layers[layer]
	nOnesRE := regexp.MustCompile("1")
	matches := nOnesRE.FindAllStringIndex(minZeroLayer, -1)
	nOnes := len(matches)

	nTwosRE := regexp.MustCompile("2")
	matches = nTwosRE.FindAllStringIndex(minZeroLayer, -1)
	nTwos := len(matches)

	return nOnes * nTwos
}
