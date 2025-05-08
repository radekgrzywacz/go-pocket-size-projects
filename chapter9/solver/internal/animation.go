package solver

import (
	"image"
	plt "image/color/palette"
	"log"

	"golang.org/x/image/draw"
)

func (s *Solver) countExplorableCells() int {
	explorableCells := 0
	bounds := s.maze.Bounds()
	width := (bounds.Dx() - 2) / 10  // Number of cells in width
	height := (bounds.Dy() - 2) / 10 // Number of cells in height
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			px := 2 + x*10
			py := 2 + y*10
			// Check the top-left pixel of the 8x8 cell
			if s.maze.RGBAAt(px, py) != s.palette.wall {
				explorableCells++
			}
		}
	}
	return explorableCells
}

func (s *Solver) registerExploredPixels() {
	const totalExpectedFrames = 30
	s.drawCurrentFrameToGIF()
	explorableCells := s.countExplorableCells()
	frameInterval := 1
	if explorableCells > totalExpectedFrames {
		frameInterval = explorableCells / totalExpectedFrames
	}
	log.Printf("explorableCells: %d, frameInterval: %d", explorableCells, frameInterval)
	cellsExplored := 0
	for {
		select {
		case <-s.quit:
			s.drawCurrentFrameToGIF()
			return
		case pos := <-s.exploredPixels:
			s.markAsExplored(pos)
			cellsExplored++
			log.Printf("Explored cell %d at %v", cellsExplored, pos)
			if cellsExplored%frameInterval == 0 || cellsExplored == explorableCells {
				s.drawCurrentFrameToGIF()
			}
		}
	}
}

func (s *Solver) drawCurrentFrameToGIF() {
	const (
		gifWidth      = 500
		frameDuration = 20
	)

	frame := image.NewPaletted(image.Rect(0, 0, gifWidth, gifWidth*s.maze.Bounds().Dy()/s.maze.Bounds().Dx()), plt.Plan9)
	draw.CatmullRom.Scale(frame, frame.Rect, s.maze, s.maze.Bounds(), draw.Over, nil)

	s.animation.Image = append(s.animation.Image, frame)
	s.animation.Delay = append(s.animation.Delay, frameDuration)
}