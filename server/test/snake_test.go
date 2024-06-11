package tests

import (
	"testing"

	"github.com/nlvphu/SnakeOnl/server/internal/game"
)

func TestNewSnake(t *testing.T) {
	snake := game.NewSnake("test")
	if snake.ID != "test" {
		t.Errorf("Expected ID to be 'test', got %s", snake.ID)
	}
	if len(snake.Body) != 1 {
		t.Errorf("Expected body length to be 1, got %d", len(snake.Body))
	}
}

func TestMoveSnake(t *testing.T) {
	snake := game.NewSnake("test")
	snake.Move()
	if snake.Body[0].X != 1 || snake.Body[0].Y != 0 {
		t.Errorf("Expected head position to be (1,0), got (%d,%d)", snake.Body[0].X, snake.Body[0].Y)
	}
}
