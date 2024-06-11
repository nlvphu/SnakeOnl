package models

type Position struct {
	X, Y int
}

type Snake struct {
	ID        string
	Body      []Position
	Direction string
}
