package game

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Player struct {
	ID   string
	Conn *websocket.Conn
}

var players = make(map[string]*Player)
var playersMutex sync.Mutex

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	playerID := r.URL.Query().Get("id")
	player := &Player{ID: playerID, Conn: conn}
	playersMutex.Lock()
	players[playerID] = player
	playersMutex.Unlock()

	snake := NewSnake(playerID)
	state.Mutex.Lock()
	state.Snakes[playerID] = snake
	state.Mutex.Unlock()

	go handleMessages(player)

	defer func() {
		playersMutex.Lock()
		delete(players, playerID)
		playersMutex.Unlock()
		state.Mutex.Lock()
		delete(state.Snakes, playerID)
		state.Mutex.Unlock()
		conn.Close()
	}()
}

func handleMessages(player *Player) {
	defer player.Conn.Close()

	for {
		_, msg, err := player.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("Read error: %v\n", err)
			}
			break
		}

		fmt.Printf("Received: %s from player %s\n", string(msg), player.ID)

		state.Mutex.Lock()
		if snake, ok := state.Snakes[player.ID]; ok {
			snake.ChangeDirection(string(msg))
		}
		state.Mutex.Unlock()
	}
}

func broadcastGameState() {
	for {
		state.Mutex.Lock()
		for _, player := range players {
			if err := player.Conn.WriteJSON(state.Snakes); err != nil {
				fmt.Printf("Write error: %v\n", err)
				player.Conn.Close()
				delete(players, player.ID)
			}
		}
		state.Mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}

func init() {
	go UpdateGameState()
	go broadcastGameState()
}
