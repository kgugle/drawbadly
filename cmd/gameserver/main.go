package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kgugle/drawbadly/pkg/engine"
)

var (
	socketEndpoint int
)

func init() {
	socketEndpoint = *flag.Int("socket-endpoint", 9000, "Client Websocket Endpoint")
}

func main() {
	flag.Parse()

	rootURL := fmt.Sprintf("localhost:%d", socketEndpoint)
	// TODO: replace with game hub
	game := engine.NewGame(rootURL)

	// define endpoints
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		engine.GameHandler(game, w, r)
	})
	http.HandleFunc("/liveness", engine.LivenessHandler)

	log.Println("GameServer running on localhost:", socketEndpoint)
	if err := http.ListenAndServe(rootURL, nil); err != nil {
		log.Fatal("game_server_error", err)
	}
}
