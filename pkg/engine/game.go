package engine

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Pixel ...
type Pixel struct {
	X     uint32
	Y     uint32
	Color uint8
}

// ConnectionStatus describes the connection status of server <-> player.
type ConnectionStatus int

// connection states
const (
	Connected ConnectionStatus = 0
	// ReadFailing is set when a single read fails on PlayerState.Conn. This
	// gives us an early indication that a client might fail.
	ReadFailing ConnectionStatus = 1
	// WriteFailing is set when a single write fails on PlayerState.Conn. This
	// gives us an early indication that a client might fail.
	WriteFailing ConnectionStatus = 2
	// Disconnected describes a player that has temporarily disappeared, and
	// is attempting to reconnect.
	Disconnected ConnectionStatus = 0
)

// PlayerState ...
type PlayerState struct {
	// basic state
	ID   uint8 // Max 256 players per game
	Conn *websocket.Conn

	// player information
	IP        string
	FirstName string

	Mutex  sync.RWMutex
	Score  []int // Append scores
	Status ConnectionStatus
}

// Game ...
type Game struct {
	PlayerMap sync.Map // ID -> PlayerState
	Drawer    *PlayerState

	Start time.Time

	PixelChan chan Pixel
}

func NewGame() *Game {
	return &Game{
		PixelChan: make(chan Pixel),
	}
}

// func (g *Game) registerPlayer(ws *websocket.Conn, drawer bool) {
// 	g.PlayerMap.Store()
// }
