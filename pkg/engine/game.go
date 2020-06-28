package engine

import (
	"log"

	"github.com/gorilla/websocket"
)

func (g *Game) BroadcastPixel(pixelData []byte) error {
	c := g.Players
	c.Lock()
	defer c.Unlock()
	for i, v := range c.internal {
		if i == g.DrawerID {
			continue
		}
		playerConn := v.Conn

		if err := playerConn.WriteMessage(websocket.BinaryMessage, pixelData); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) logGameState() error {
	log.Printf("------------GAME %s: \n", g.ID)
	c := g.Players
	c.Lock()
	defer c.Unlock()
	for i, v := range c.internal {
		if i == g.DrawerID {
			log.Printf("\t(DRAWER)")
		}
		log.Printf("\tNAME: %s", v.Name)
		log.Printf("\tID: %d", v.ID)
		log.Printf("\tSCORE: %d", v.Score)
		log.Printf("\tCONNECTION STATUS: %d", v.Status)
		log.Println()
	}
	return nil
}
