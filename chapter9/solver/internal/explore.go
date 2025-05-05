package solver

import (
	"image"
	"image/color"
	"log"
)

func (s *Solver) listenToBranches() {
	for p := range s.pathsToExplore {
		go s.explore(p)
		if s.solutionFound() {
			return
		}
	}
}

func (s *Solver) solutionFound() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.solution != nil
}

func (s *Solver) explore(pathToBranch *path) {
	log.Printf("In explore")
	if pathToBranch == nil {
		return
	}

	pos := pathToBranch.at

	for !s.solutionFound() {
		candidates := make([]cell, 0, 3)
		for _, n := range neighbours(pos) {
			if pathToBranch.isPreviousStep(n) {
				continue
			}

			from := cellToPixel(pos.X, pos.Y)
			to := cellToPixel(n.X, n.Y)

			if isConnected(from, to, s.maze, s.palette.path) {
				if isTreasure(n, s.maze, s.palette.treasure) {
					s.mutex.Lock()
					defer s.mutex.Unlock()
					s.solution = &path{previousStep: pathToBranch, at: n}
					log.Printf("Treasure found at %v", pos)
					return
				}
				candidates = append(candidates, n)
			}

		}

		if len(candidates) == 0 {
			log.Printf("I must have taken the wrong turn at position %v", pos)
			return
		}

		for _, candidate := range candidates[1:] {
			branch := &path{previousStep: pathToBranch, at: candidate}
			s.pathsToExplore <- branch
		}

		pathToBranch = &path{previousStep: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}
}

func (p path) isPreviousStep(n cell) bool {
	return p.previousStep != nil && p.previousStep.at == n
}

func isConnected(a, b image.Point, maze *image.RGBA, pathColor color.RGBA) bool {
	mid := image.Point{
		X: (a.X + b.X) / 2,
		Y: (a.Y + b.Y) / 2,
	}
	return maze.RGBAAt(mid.X, mid.Y) == pathColor
}


func isTreasure(c cell, maze *image.RGBA, treasureColor color.RGBA) bool {
	p := cellToPixel(c.X, c.Y)
	return maze.RGBAAt(p.X, p.Y) == treasureColor
}
