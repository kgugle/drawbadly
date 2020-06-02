package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

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

	// TODO: replace with game hub
	game := engine.NewGame()

	// define endpoints
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		engine.GameHandler(game, w, r)
	})
	http.HandleFunc("/liveness", engine.LivenessHandler)

	log.Println("GameServer running on localhost:", socketEndpoint)
	if err := http.ListenAndServe("localhost:"+strconv.Itoa(socketEndpoint), nil); err != nil {
		log.Fatal("game_server_error", err)
	}
}
