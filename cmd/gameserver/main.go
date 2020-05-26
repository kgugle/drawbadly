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
	http.HandleFunc("/pixel", engine.HandlePixelStream)
	http.HandleFunc("/liveness", liveness)

	log.Println("GameServer running on localhost:", socketEndpoint)
	if err := http.ListenAndServe("localhost:"+strconv.Itoa(socketEndpoint), nil); err != nil {
		log.Fatal("game_server_error", err)
	}
}

// liveness is a sanity check to ensure the server is running
func liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
