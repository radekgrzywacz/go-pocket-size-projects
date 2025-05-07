package solver

import (
	"image"
	"image/color"
)

type palette struct {
	wall     color.RGBA
	path     color.RGBA
	entrance color.RGBA
	treasure color.RGBA
	solution color.RGBA
}

func defaultPalette() palette {
	return palette{
		wall:     color.RGBA{R: 0, G: 0, B: 0, A: 255},
		path:     color.RGBA{R: 255, G: 255, B: 255, A: 255},
		entrance: color.RGBA{R: 0, G: 191, B: 255, A: 255},
		treasure: color.RGBA{R: 255, G: 0, B: 128, A: 255},
		solution: color.RGBA{R: 255, G: 0, B: 0, A: 255},
	}
}

func setCell(cellX int, cellY int, img *image.RGBA, rgba color.RGBA) {
	pixelX := 2 + cellX*10 // convert cell to pixel
	pixelY := 2 + cellY*10 // convert cell to pixel

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			img.Set(pixelX+i, pixelY+j, rgba)
		}
	}
}
