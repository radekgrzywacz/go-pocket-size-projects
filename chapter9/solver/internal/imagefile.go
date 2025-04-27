package solver 

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func openMaze(imagePath string) (*image.RGBA, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("Unable to open image %s: %w", imagePath, err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("Unable to load input image from %s: %w", img, err)
	}

	rgbaImage, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("Expected RGBA image, got %T", img)
	}

	return rgbaImage, nil
}

func (s *Solver) SaveSolution(outputPath string) error {
	return nil
}
