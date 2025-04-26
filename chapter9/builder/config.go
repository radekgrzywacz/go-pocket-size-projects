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
		wallColor:     color.RGBA{0, 0, 0, 255},
		pathColor:     color.RGBA{255, 255, 255, 255},
		entranceColor: color.RGBA{0, 191, 255, 255},
		treasureColor: color.RGBA{255, 0, 128, 255},
	}
}
