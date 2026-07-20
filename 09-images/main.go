// Video explanation: https://www.youtube.com/watch?v=P3hUrm1-HeI&t=12s

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
	pixel         func(int, int) uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	v := i.pixel(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{
		width:  256,
		height: 128,
		pixel: func(x, y int) uint8 {
			return uint8(x*x + y*y)
		},
	}
	pic.ShowImage(m)
}
