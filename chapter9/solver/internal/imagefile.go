package solver

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
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
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("Unable to create output image file at %s", outputPath)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("Unable to close file"))
		}
	}()

	stepsFromTreasure := s.solution
	for stepsFromTreasure != nil {
		pixel := cellToPixel(stepsFromTreasure.at.X, stepsFromTreasure.at.Y)
		s.maze.Set(pixel.X, pixel.Y, s.palette.solution)
		stepsFromTreasure = stepsFromTreasure.previousStep
	}

	err = png.Encode(f, s.maze)
	if err != nil {
		return fmt.Errorf("Unable to write output image at %s: %w", outputPath, err)
	}
	log.Print("Saved")
	return nil
}
