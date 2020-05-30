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

	// init game state

	game := engine.NewGame()

	// define endpoints
	http.HandleFunc("/drawer", func(w http.ResponseWriter, r *http.Request) {
		engine.DrawerHandler(game, w, r)
	})
	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		engine.PlayerHandler(game, w, r)
	})
	http.HandleFunc("/liveness", engine.LivenessHandler)

	log.Println("GameServer running on localhost:", socketEndpoint)
	if err := http.ListenAndServe("localhost:"+strconv.Itoa(socketEndpoint), nil); err != nil {
		log.Fatal("game_server_error", err)
	}
}
