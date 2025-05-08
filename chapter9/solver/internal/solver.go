package solver

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"sync"
)

type Solver struct {
	maze           *image.RGBA
	palette        palette
	pathsToExplore chan *path
	solution       *path
	mutex          sync.Mutex
	quit           chan struct{}
	exploredPixels chan cell
	animation      *gif.GIF
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
		quit:           make(chan struct{}),
		exploredPixels: make(chan cell),
		animation:      &gif.GIF{},
	}, nil
}

func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("Unable to find entrance: %w", err)
	}

	log.Printf("starting at %v", entrance)
	s.pathsToExplore <- &path{previousStep: nil, at: cell{entrance.X, entrance.Y}}
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		s.registerExploredPixels()
	}()

	go func() {
		defer wg.Done()
		s.listenToBranches()
	}()

	wg.Wait()

	s.writeLastFrame()

	return nil
}

func (s *Solver) writeLastFrame() {
	stepsFromTreasure := s.solution
	for stepsFromTreasure != nil && stepsFromTreasure.previousStep != nil {
		from := cellToPixel(stepsFromTreasure.at.X, stepsFromTreasure.at.Y)
		to := cellToPixel(stepsFromTreasure.previousStep.at.X, stepsFromTreasure.previousStep.at.Y)
		drawLine(s.maze, from, to, s.palette.solution)
		stepsFromTreasure = stepsFromTreasure.previousStep
	}
	const solutionFrameDuration = 300 // 3 seconds
	s.drawCurrentFrameToGIF()
	s.animation.Delay[len(s.animation.Delay)-1] = solutionFrameDuration
	log.Printf("Added final solution frame")
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

func (s *Solver) markAsExplored(pos cell) {
	px := 2 + pos.X*10
	py := 2 + pos.Y*10
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			s.maze.Set(px+i, py+j, s.palette.explored)
		}
	}
}
