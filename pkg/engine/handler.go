package engine

import (
	"encoding/binary"
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
func DrawerHandler(w http.ResponseWriter, r *http.Request) {
	drawConn, err := upgrade_connection_to_ws(w, r)
	if err != nil {
		log.Print("draw_stream_upgrade_failed", err)
		return
	}
	log.Println("drawer connected")
	defer drawConn.Close()

	for {
		_, drawBuffer, err := drawConn.ReadMessage()
		if err != nil {
			log.Println("draw_stream_read_error", err)
			break
		}

		// TODO: let's marshal this into a struct
		x := binary.LittleEndian.Uint32(drawBuffer[:4])
		y := binary.LittleEndian.Uint32(drawBuffer[4:])
		log.Printf("(%d,%d)", x, y)
	}

}

// PlayerHandler receives the x,y coordinates sent by the drawing client.
func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	playerConn, err := upgrade_connection_to_ws(w, r)
	if err != nil {
		log.Print("player_stream_upgrade_failed", err)
		return
	}
	log.Println("player connected")
	defer playerConn.Close()

	for {
		if err := playerConn.WriteMessage(websocket.TextMessage, []byte("test\n")); err != nil {
			log.Println("player_stream_write_error", err)
			return
		}
		time.Sleep(2 * time.Second)
	}
}
