package engine

import (
	"github.com/gorilla/websocket"
)

type GameState struct {
	drawer_conn *websocket.Conn
	player_conn *websocket.Conn
	pixels      chan Pixel
}

type Pixel struct {
	x, y uint32
}

func InitGameState() *GameState {
	return &GameState{pixels: make(chan Pixel)}
}
