package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run . <width> <height>")
	}
	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid width:", err)
	}
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid height:", err)
	}

	maze := createMaze(width, height)

	file, err := os.Create(fmt.Sprintf("../solver/mazes/maze%d_%d.png", width, height))
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}
	defer file.Close()
	if err := png.Encode(file, &maze); err != nil {
		log.Fatal("Failed to encode PNG:", err)
	}
}

type cell struct {
	x       int
	y       int
	visited bool
}

func createMaze(width int, height int) image.RGBA {
	colors := defaultColors()
	imgWidth := 8*width + 2*width + 2
	imgHeight := 8*height + 2*height + 2
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	cells := make([][]cell, height)

	for y := range cells {
		cells[y] = make([]cell, width)
		for x := range cells[y] {
			cells[y][x] = cell{x, y, false}
		}
	}

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			img.Set(x, y, colors.wallColor)
		}
	}

	entrance := cell{0, height / 2, true}
	if entrance.y > height {
		log.Fatalf("Wrong entrance")
	}
	cells[entrance.y][entrance.x] = entrance
	colorEntrance(entrance, img, colors)

	treasurePlaced := false

	generateMaze(entrance.x, entrance.y, width, height, cells, img, &treasurePlaced)

	return *img
}

func generateMaze(x, y int, width int, height int, cells [][]cell, img *image.RGBA, treasurePlaced *bool) {
	colors := defaultColors()
	cells[y][x].visited = true
	setCell(x, y, img, colors.pathColor)

	if x == width-1 && !*treasurePlaced {
		colorFinish(x, y, img, colors, width)
		*treasurePlaced = true
		return
	}

	directions := [][2]int{
		{-1, 0}, // left
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
	}

	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]

		if newX >= 0 && newX < width && newY >= 0 && newY < height && !cells[newY][newX].visited {
			carvePath(x, y, newX, newY, img, colors.pathColor)
			generateMaze(newX, newY, width, height, cells, img, treasurePlaced)
		}
	}
}

func carvePath(x int, y int, x2 int, y2 int, img *image.RGBA, pathColor color.RGBA) {
	px1 := 2 + x*10
	py1 := 2 + y*10

	switch {
	case x == x2 && y == y2-1: // wall below
		for i := 0; i < 8; i++ {
			img.Set(px1+i, py1+8, pathColor)
			img.Set(px1+i, py1+9, pathColor)
		}
	case x == x2 && y == y2+1:
		for i := 0; i < 8; i++ {
			img.Set(px1+i, py1-1, pathColor)
			img.Set(px1+i, py1-2, pathColor)
		}
	case x == x2-1 && y == y2:
		for i := 0; i < 8; i++ {
			img.Set(px1+8, py1+i, pathColor)
			img.Set(px1+9, py1+i, pathColor)
		}
	case x == x2+1 && y == y2:
		for i := 0; i < 8; i++ {
			img.Set(px1-1, py1+i, pathColor)
			img.Set(px1-2, py1+i, pathColor)
		}
	}
}

func colorEntrance(entrance cell, img *image.RGBA, colors config) {
	setCell(entrance.x, entrance.y, img, colors.entranceColor)
	py := 2 + entrance.y*10
	for i := 0; i < 8; i++ {
		img.Set(0, py+i, colors.entranceColor)
		img.Set(1, py+i, colors.entranceColor)
		img.Set(2, py+i, colors.entranceColor)
	}
}

func colorFinish(x, y int, img *image.RGBA, colors config, width int) {
	py := 2 + y*10
	imgWidth := 8*width + 2*width + 2
	for i := 0; i < 8; i++ {
		img.Set(imgWidth-1, py+i, colors.treasureColor)
		img.Set(imgWidth-2, py+i, colors.treasureColor)
		img.Set(imgWidth, py+i, colors.treasureColor)
	}
}

func setCell(x, y int, img *image.RGBA, rgba color.RGBA) {
	px := 2 + x*10
	py := 2 + y*10
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			img.Set(px+i, py+j, rgba)
		}
	}
}
