package engine

import (
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

// GameHandler creates a game with a new random ID
func CreateGameHandler(gh *GameHub, w http.ResponseWriter, r *http.Request) {
	gh.registerGame(r.Host)
}

// GameHandler serves the game
func GameHandler(gh *GameHub, w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "404, no game id provided.", http.StatusNotFound)
		return
	}
	gameID := keys[0]
	g, ok := gh.Games.Load(gameID)
	if !ok {
		http.Error(w, "404, game not found.", http.StatusNotFound)
		return
	}

	playerConn, err := upgrade_connection_to_ws(w, r)
	if err != nil {
		log.Print("player_stream_upgrade_failed", err)
		return
	}
	defer playerConn.Close()

	playerID := g.registerPlayer(playerConn)
	log.Printf("player ID %d registered", playerID)
	g.logGameState()

	if g.Players.Length() == 1 {
		if ok := g.setDrawer(playerID); !ok {
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
			// x := binary.LittleEndian.Uint32(drawBuffer[:4])
			// y := binary.LittleEndian.Uint32(drawBuffer[4:8])
			// ss := binary.LittleEndian.Uint32(drawBuffer[8:])

			// fmt.Printf("send %d %d\n", x, y)
			if err := g.BroadcastPixel(drawBuffer); err != nil {
				log.Println("broadcast_pixel_error", err)
			}
		}
	} else {
		for {
		}
	}
}
