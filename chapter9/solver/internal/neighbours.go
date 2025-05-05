package solver

import "image"

func neighbours(p cell) [4]cell {
	return [...]cell{
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
	}
}

func cellToPixel(x, y int) image.Point {
	return image.Point{
		X: 2 + x*10 + 4,
		Y: 2 + y*10 + 4,
	}
}
