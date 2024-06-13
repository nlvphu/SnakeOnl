package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Capture interrupt signals to allow for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Set up WebSocket connection URL
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws", RawQuery: "id=testClient"}
	log.Printf("Connecting to %s", u.String())

	// Establish the WebSocket connection
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// Goroutine to handle incoming messages
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Read:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Send some test messages (direction commands) to the server
	directions := []string{"UP", "DOWN", "LEFT", "RIGHT"}
	for _, dir := range directions {
		log.Printf("Sending direction: %s", dir)
		err := c.WriteMessage(websocket.TextMessage, []byte(dir))
		if err != nil {
			log.Println("Write:", err)
			return
		}
		time.Sleep(time.Second)
	}

	// Wait for interrupt signal to gracefully shut down the client
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Interrupt received, shutting down...")

			// Send close message to server
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
