package solver

type path struct {
	previousStep *path
	at           cell
}

type cell struct {
	X int
	Y int
}
