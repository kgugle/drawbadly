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

type PlayerMap struct {
	sync.RWMutex
	internal map[int]*PlayerState
}

func makePlayerMap() *PlayerMap {
	return &PlayerMap{internal: make(map[int]*PlayerState)}
}

func (c *PlayerMap) LoadOrStore(key int, value *PlayerState) (actual *PlayerState, loaded bool) {
	c.Lock()
	defer c.Unlock()
	actual, ok := c.internal[key]
	if !ok {
		c.internal[key] = value
		return value, false
	}
	return actual, true
}

func (c *PlayerMap) Length() int {
	c.Lock()
	defer c.Unlock()
	return len(c.internal)
}

func (c *PlayerMap) Delete(key int) {
	c.Lock()
	defer c.Unlock()
	delete(c.internal, key)
}

// Game ...
type Game struct {
	PlayersByID *PlayerMap // ID -> PlayerState
	Drawer      *PlayerState

	Start time.Time

	PixelChan chan Pixel
}

func NewGame() *Game {
	return &Game{
		PlayersByID: makePlayerMap(),
		PixelChan:   make(chan Pixel),
	}
}

func (g *Game) registerPlayer(ws *websocket.Conn) (newID int) {
	newID = g.PlayersByID.Length()
	g.PlayersByID.LoadOrStore(newID, &PlayerState{Conn: ws})
	return newID
}
