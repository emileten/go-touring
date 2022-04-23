package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	// number of columns and rows respectively
	dx, dy int
}

func (I Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (I Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, I.dx, I.dy)
}

func (I Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x ^ y), uint8(x ^ y), uint8(255), uint8(255)}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
