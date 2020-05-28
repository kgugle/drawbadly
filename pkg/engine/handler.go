package engine

import (
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func upgrade_connection_to_ws(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var upgrader = websocket.Upgrader{
		HandshakeTimeout: 3 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// EnableCompression: true,
	}
	return upgrader.Upgrade(w, r, nil)
}

// LivenessHandler is a sanity check to ensure the server is running
func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// DrawerHandler receives the x,y coordinates sent by the drawing client.
func DrawerHandler(gs *GameState, w http.ResponseWriter, r *http.Request) {
	drawConn, err := upgrade_connection_to_ws(w, r)
	if err != nil {
		log.Print("draw_stream_upgrade_failed", err)
		return
	}
	log.Println("drawer connected")
	defer drawConn.Close()

	gs.drawer_conn = drawConn

	for {
		_, drawBuffer, err := drawConn.ReadMessage()
		if err != nil {
			log.Println("draw_stream_read_error", err)
			break
		}

		// TODO: let's marshal this into a struct
		x := binary.LittleEndian.Uint32(drawBuffer[:4])
		y := binary.LittleEndian.Uint32(drawBuffer[4:])
		gs.pixels <- Pixel{x, y}
		log.Printf("server received (%d,%d)", x, y)
	}
}

// PlayerHandler receives the x,y coordinates sent by the drawing client.
func PlayerHandler(gs *GameState, w http.ResponseWriter, r *http.Request) {
	playerConn, err := upgrade_connection_to_ws(w, r)
	if err != nil {
		log.Print("player_stream_upgrade_failed", err)
		return
	}
	log.Println("player connected")
	defer playerConn.Close()

	gs.player_conn = playerConn

	for {
		select {
		case pixel, ok := <-gs.pixels:
			if !ok {
				log.Println("pixel_channel_error", err)
				return
			}
			msg := fmt.Sprintf("%d %d\n", pixel.x, pixel.y)
			if err := playerConn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("player_stream_write_error", err)
				return
			}
		}
	}
}
