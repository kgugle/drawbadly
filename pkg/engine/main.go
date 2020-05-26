package engine

import (
	"encoding/binary"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// HandlePixelStream prints the x,y coordinates sent by the drawing client.
func HandlePixelStream(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		HandshakeTimeout: 3 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// EnableCompression: true,
	}
	pixelConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("pixel_stream_upgrade_failed", err)
		return
	}
	defer pixelConn.Close()

	for {
		_, pixelBuffer, err := pixelConn.ReadMessage()
		if err != nil {
			log.Println("pixel_stream_read_error", err)
			break
		}

		// TODO: let's marshal this into a struct
		x := binary.LittleEndian.Uint32(pixelBuffer[:4])
		y := binary.LittleEndian.Uint32(pixelBuffer[4:])
		log.Printf("(%d,%d)", x, y)
	}

}
