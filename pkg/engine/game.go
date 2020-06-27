package engine

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// ConnectionStatus describes the connection status of server <-> player.
type ConnectionStatus int

const (
	Connected ConnectionStatus = iota
	// ReadFailing is set when a single read fails on Player.Conn. This
	// gives us an early indication that a client might fail.
	ReadFailing
	// WriteFailing is set when a single write fails on Player.Conn. This
	// gives us an early indication that a client might fail.
	WriteFailing
	// Disconnected describes a player that has temporarily disappeared, and
	// is attempting to reconnect.
	Disconnected
)

// Player ...
type Player struct {
	ID     uint8 // Max 256 players per game
	Conn   *websocket.Conn
	Score  int
	IP     string
	Name   string
	Mutex  sync.RWMutex
	Status ConnectionStatus
}

type PlayerMap struct {
	sync.RWMutex
	internal map[int]*Player
}

func makePlayerMap() *PlayerMap {
	return &PlayerMap{internal: make(map[int]*Player)}
}

func (c *PlayerMap) Load(key int) (value *Player, ok bool) {
	c.RLock()
	defer c.RUnlock()
	result, ok := c.internal[key]
	return result, ok
}

func (c *PlayerMap) Store(key int, value *Player) {
	c.Lock()
	defer c.Unlock()
	c.internal[key] = value
}

func (c *PlayerMap) LoadOrStore(key int, value *Player) (actual *Player, loaded bool) {
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
	Players  *PlayerMap // Player ID -> Player
	DrawerID int

	ID      string
	RootURL string
}

func NewGame(root_url string) *Game {
	// TODO: randomly generate an ID
	return &Game{
		ID:      "asdf",
		Players: makePlayerMap(),
		RootURL: root_url,
	}
}

func (g *Game) setDrawer(playerID int) (ok bool) {
	if _, ok := g.Players.Load(playerID); !ok {
		return false
	}
	g.DrawerID = playerID
	return true
}

func (g *Game) Drawer() *Player {
	drawer, _ := g.Players.Load(g.DrawerID)
	return drawer
}

func (g *Game) gameURL() (url string) {
	return fmt.Sprintf("%s/%s", g.RootURL, g.ID)
}

func (g *Game) registerPlayer(ws *websocket.Conn) (newID int) {
	newPlayer := &Player{
		Conn: ws,
	}
	newID = g.Players.Length()
	g.Players.Store(newID, newPlayer)
	return newID
}

func (g *Game) BroadcastPixel(pixelData []byte) error {
	c := g.Players
	c.Lock()
	defer c.Unlock()
	for i, v := range c.internal {
		if i == g.DrawerID {
			continue
		}
		playerConn := v.Conn

		if err := playerConn.WriteMessage(websocket.BinaryMessage, pixelData); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) logGameState() error {
	log.Printf("------------GAME %s: \n", g.ID)
	c := g.Players
	c.Lock()
	defer c.Unlock()
	for i, v := range c.internal {
		if i == g.DrawerID {
			log.Printf("\t(DRAWER)")
		}
		log.Printf("\tNAME: %s", v.Name)
		log.Printf("\tID: %d", v.ID)
		log.Printf("\tSCORE: %d", v.Score)
		log.Printf("\tCONNECTION STATUS: %d", v.Status)
		log.Println()
	}
	return nil
}
