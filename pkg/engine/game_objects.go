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

type GameStatus int

const (
	GameWaiting GameStatus = iota
	GameStarted
	GameEnded
)

// Game ...
type Game struct {
	Players  *PlayerMap // Player ID -> Player
	DrawerID int
	Status   GameStatus
	ID       string
	URL      string
}

func NewGame(ID, rootURL string) *Game {
	log.Printf("game ID %s created", ID)
	return &Game{
		ID:      ID,
		Players: makePlayerMap(),
		URL:     fmt.Sprintf("%s/%s", rootURL, ID),
		Status:  GameWaiting,
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

func (g *Game) registerPlayer(ws *websocket.Conn) (newID int) {
	newPlayer := &Player{
		Conn: ws,
	}
	newID = g.Players.Length()
	g.Players.Store(newID, newPlayer)
	return newID
}

type GameHub struct {
	Games *GameMap // Game ID -> Game
}

func NewGameHub() *GameHub {
	return &GameHub{
		makeGameMap(),
	}
}

func (gh *GameHub) registerGame(hostname string) (newID string) {
	newID = GenerateID(8)
	newGame := NewGame(newID, hostname)
	_, prs := gh.Games.LoadOrStore(newID, newGame)
	for prs { // regenerate if we get a duplicate
		newID = GenerateID(8)
		newGame.ID = newID
		_, prs = gh.Games.LoadOrStore(newID, newGame)
	}

	return newID
}
