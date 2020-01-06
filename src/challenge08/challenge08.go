package challenge08

import (
	"fmt"
	"regexp"
)

// Image represents an Image
type Image struct {
	data        string
	height      int
	width       int
	renderLayer string
	layers      []string
}

// NewImage creates an image
func NewImage(data string, height, width int) *Image {
	var layers []string
	layerLength := height * width

	var dataLen = len(data)
	var layerStart = 0
	for nLayers := 0; nLayers*layerLength < dataLen; nLayers++ {
		layers = append(layers, data[layerStart:layerStart+layerLength])
		layerStart += layerLength
	}

	var renderLayer string

	var image = Image{
		data,
		height,
		width,
		renderLayer,
		layers,
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			renderLayer = renderLayer + image.visibileElementAt(x, y)
		}
	}
	image.renderLayer = renderLayer
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

// GetRender returns the rendered layer
func (img *Image) GetRender() string {
	return img.renderLayer
}

const (
	black       = '0'
	white       = '1'
	transparent = '2'
)

func (img *Image) visibileElementAt(x, y int) string {
	for _, layer := range img.layers {
		var rows []string

		for i := 0; i < img.height; i++ {
			start := i * img.width
			rows = append(rows, layer[start:start+img.width])
		}

		char := rows[y][x]
		switch char {
		case black:
			return "0"
		case white:
			return "1"
		case transparent:
		}
	}
	return " "
}

// Print Image
func (img *Image) Print() {
	for i := 0; i < img.height; i++ {
		start := i * img.width
		row := img.renderLayer[start : start+img.width]
		fmt.Printf("'%s'\n", row)
	}
}
