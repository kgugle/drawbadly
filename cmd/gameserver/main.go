package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/kgugle/drawbadly/pkg/engine"
)

var (
	socketPort int
)

func init() {
	flag.IntVar(&socketPort, "port", 9000, "Client Websocket Endpoint")
}

func main() {
	flag.Parse()
	rootURL := fmt.Sprintf("localhost:%d", socketPort)

	rand.Seed(time.Now().UnixNano())

	gameHub := engine.NewGameHub()

	// define endpoints
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		engine.CreateGameHandler(gameHub, w, r)
	})
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		engine.GameHandler(gameHub, w, r)
	})
	http.HandleFunc("/liveness", engine.LivenessHandler)

	log.Printf("GameServer running on %s\n", rootURL)
	if err := http.ListenAndServe(rootURL, nil); err != nil {
		log.Fatal("game_server_error", err)
	}
}
