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

	// define endpoints
	http.HandleFunc("/drawer", engine.DrawerHandler)
	http.HandleFunc("/player", engine.PlayerHandler)
	http.HandleFunc("/liveness", engine.LivenessHandler)

	log.Println("GameServer running on localhost:", socketEndpoint)
	if err := http.ListenAndServe("localhost:"+strconv.Itoa(socketEndpoint), nil); err != nil {
		log.Fatal("game_server_error", err)
	}
}
