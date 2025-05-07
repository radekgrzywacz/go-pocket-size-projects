package solver

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
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
	for stepsFromTreasure != nil && stepsFromTreasure.previousStep != nil {
		from := cellToPixel(stepsFromTreasure.at.X, stepsFromTreasure.at.Y)
		to := cellToPixel(stepsFromTreasure.previousStep.at.X, stepsFromTreasure.previousStep.at.Y)
	
		drawLine(s.maze, from, to, s.palette.solution)
		stepsFromTreasure = stepsFromTreasure.previousStep
	}

	err = png.Encode(f, s.maze)
	if err != nil {
		return fmt.Errorf("Unable to write output image at %s: %w", outputPath, err)
	}
	log.Print("Saved")
	return nil
}

func drawLine(img *image.RGBA, from, to image.Point, clr color.Color) {
	dx := to.X - from.X
	dy := to.Y - from.Y
	steps := int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))

	if steps == 0 {
		img.Set(from.X, from.Y, clr)
		return
	}

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		x := int(math.Round(float64(from.X) + t*float64(dx)))
		y := int(math.Round(float64(from.Y) + t*float64(dy)))
		img.Set(x, y, clr)
	}
}
