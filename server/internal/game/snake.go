package game

import (
	"sync"
	"time"
)

type Position struct {
	X, Y int
}

type Snake struct {
	ID        string
	Body      []Position
	Direction string
	Mutex     sync.Mutex
}

type GameState struct {
	Snakes map[string]*Snake
	Mutex  sync.Mutex
}

var state = GameState{
	Snakes: make(map[string]*Snake),
}

func NewSnake(id string) *Snake {
	return &Snake{
		ID:        id,
		Body:      []Position{{X: 0, Y: 0}},
		Direction: "RIGHT",
	}
}

func (s *Snake) Move() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	head := s.Body[0]
	newHead := head

	switch s.Direction {
	case "UP":
		newHead.Y -= 1
	case "DOWN":
		newHead.Y += 1
	case "LEFT":
		newHead.X -= 1
	case "RIGHT":
		newHead.X += 1
	}

	s.Body = append([]Position{newHead}, s.Body[:len(s.Body)-1]...)
}

func UpdateGameState() {
	for {
		time.Sleep(500 * time.Millisecond)
		state.Mutex.Lock()
		for _, snake := range state.Snakes {
			snake.Move()
		}
		state.Mutex.Unlock()
	}
}

func (s *Snake) ChangeDirection(newDirection string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Direction = newDirection
}
