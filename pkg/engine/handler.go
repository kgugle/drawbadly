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

// GameHandler serves the game
func GameHandler(g *Game, w http.ResponseWriter, r *http.Request) {
	playerConn, err := upgrade_connection_to_ws(w, r)
	if err != nil {
		log.Print("player_stream_upgrade_failed", err)
		return
	}
	log.Println("player connected")
	defer playerConn.Close()

	playerID := g.registerPlayer(playerConn)
	log.Printf("player registered with ID %d", playerID)

	if g.PlayersByID.Length() == 1 {
		if ok := g.setDrawerID(playerID); !ok {
			log.Print("set_drawer_state_failed")
			return
		}

		for {
			_, drawBuffer, err := playerConn.ReadMessage()

			if err != nil {
				log.Println("draw_stream_read_error", err)
				break
			}

			// TODO: let's marshal this into a struct
			x := binary.LittleEndian.Uint32(drawBuffer[:4])
			y := binary.LittleEndian.Uint32(drawBuffer[4:8])
			ss := binary.LittleEndian.Uint32(drawBuffer[8:])

			g.PlayersByID.BroadcastPixel(Pixel{x, y, uint8(0), ss != 0})
		}
	} else {
		player, _ := g.PlayersByID.Load(playerID)
		for {
			select {
			case pixel, ok := <-player.PixelChan:
				if !ok {
					log.Println("pixel_channel_error")
					return
				}
				var ss uint32
				if pixel.StrokeStart {
					ss = 1
				} else {
					ss = 0
				}
				msg := fmt.Sprintf("%d %d %d\n", pixel.X, pixel.Y, ss)

				if err := playerConn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
					log.Println("player_stream_write_error", err)
					return
				}
			}
		}
	}
}
