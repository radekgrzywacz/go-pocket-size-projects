package main

import "image/color"

type config struct {
	wallColor     color.RGBA
	pathColor     color.RGBA
	entranceColor color.RGBA
	treasureColor color.RGBA
}

func defaultColors() config {
	return config{
		wallColor:     color.RGBA{R: 0, G: 0, B: 0, A: 255},
		pathColor:     color.RGBA{R: 255, G: 255, B: 255, A: 255},
		entranceColor: color.RGBA{R: 0, G: 191, B: 255, A: 255},
		treasureColor: color.RGBA{R: 255, G: 0, B: 128, A: 255},
	}
}
