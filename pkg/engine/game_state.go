package engine

// import (
// 	"encoding/binary"
// 	"log"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// func readPixels() {
// 	for {
// 		_, drawBuffer, err := drawConn.ReadMessage()
// 		if err != nil {
// 			log.Println("draw_stream_read_error", err)
// 			break
// 		}

// 		// TODO: let's marshal this into a struct
// 		x := binary.LittleEndian.Uint32(drawBuffer[:4])
// 		y := binary.LittleEndian.Uint32(drawBuffer[4:])
// 		log.Printf("(%d,%d)", x, y)
// 	}

// }
