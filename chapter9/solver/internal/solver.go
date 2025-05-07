package solver

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"sync"
)

type Solver struct {
	maze           *image.RGBA
	palette        palette
	pathsToExplore chan *path
	solution       *path
	mutex          sync.Mutex
}

func New(imagePath string) (*Solver, error) {
	img, err := openMaze(imagePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot open maze image: %w", err)
	}

	return &Solver{
		maze:           img,
		palette:        defaultPalette(),
		pathsToExplore: make(chan *path, 1),
	}, nil
}

func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("Unable to find entrance: %w", err)
	}

	log.Printf("starting at %v", entrance)
	s.pathsToExplore <- &path{previousStep: nil, at: cell{entrance.X, entrance.Y}}

	s.listenToBranches()
	file, err := os.Create("output.png")
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}
	defer file.Close()
	if err := png.Encode(file, s.maze); err != nil {
		log.Fatal("Failed to encode PNG:", err)
	}

	return nil
}

func (s *Solver) findEntrance() (image.Point, error) {
	for row := s.maze.Bounds().Min.Y; row < s.maze.Bounds().Max.Y; row++ {
		for col := s.maze.Bounds().Min.X; col < s.maze.Bounds().Max.X; col++ {
			if s.maze.RGBAAt(col, row) == s.palette.entrance {
				cellX := (col - 2) / 10
				cellY := (row - 2) / 10
				return image.Point{X: cellX, Y: cellY}, nil
			}
		}
	}

	return image.Point{}, fmt.Errorf("entrance position not found")
}
