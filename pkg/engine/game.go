package engine

import (
	"fmt"
	// "log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

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

func (c *PlayerMap) BroadcastPixel(pixelData []byte) error {
	c.Lock()
	defer c.Unlock()
	for _, v := range c.internal {
		playerConn := v.Conn

		if err := playerConn.WriteMessage(websocket.BinaryMessage, pixelData); err != nil {
			return err
		}
	}
	return nil
}

func (c *PlayerMap) Load(key int) (value *PlayerState, ok bool) {
	c.RLock()
	defer c.RUnlock()
	result, ok := c.internal[key]
	return result, ok
}

func (c *PlayerMap) Store(key int, value *PlayerState) {
	c.Lock()
	defer c.Unlock()
	c.internal[key] = value
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
	c.RLock()
	defer c.RUnlock()
	return len(c.internal)
}

func (c *PlayerMap) Delete(key int) {
	c.Lock()
	defer c.Unlock()
	delete(c.internal, key)
}

// Game ...
type Game struct {
	PlayersByID *PlayerMap // Player ID -> PlayerState
	DrawerID    int

	ID    string
	Start time.Time
}

func NewGame() *Game {
	// TODO: randomly generate an ID
	return &Game{
		ID:          "asdf",
		PlayersByID: makePlayerMap(),
	}
}

func (g *Game) setDrawerID(playerID int) (ok bool) {
	if _, ok := g.PlayersByID.Load(playerID); !ok {
		return false
	}
	g.DrawerID = playerID
	return true
}

func (g *Game) getDrawer() *PlayerState {
	drawer, _ := g.PlayersByID.Load(g.DrawerID)
	return drawer
}

func (g *Game) gameURL() (url string) {
	return fmt.Sprintf("localhost:9000/%s", g.ID)
}

func (g *Game) registerPlayer(ws *websocket.Conn) (newID int) {
	newPlayer := &PlayerState{
		Conn: ws,
	}
	newID = g.PlayersByID.Length()
	g.PlayersByID.Store(newID, newPlayer)
	return newID
}
