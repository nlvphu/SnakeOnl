package main

import (
	"fmt"
	"net/http"

	"github.com/nlvphu/SnakeOnl/server/internal/game"
)

func main() {
	http.HandleFunc("/ws", game.HandleConnections)
	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
