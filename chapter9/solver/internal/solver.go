package solver

import (
	"fmt"
	"image"
)

type Solver struct {
	maze *image.RGBA
}

func New(imagePath string) (*Solver, error) {
	img, err := openMaze(imagePath)
	if err != nil {
	  return nil, fmt.Errorf("Cannot open maze image: %w", err)
	}

	return &Solver{
		maze: img,
	}, nil
}

func (s *Solver) Solve() error {
	return nil
}