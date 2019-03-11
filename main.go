package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tronbattle/server/game"
	"github.com/tronbattle/server/socket"
)

const addr = ":8080"

var upgrader = websocket.Upgrader{
	// allow all connections by default
	CheckOrigin: func(r *http.Request) bool { return true },
}

func socketHandler(gameServer *game.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start conn")

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("upgrade error: %s", err)
			return
		}
		defer c.Close()

		conn := socket.NewConnection(c, gameServer)

		conn.ListenInputEvents()
	}
}

func main() {
	gameMap := game.NewMap()
	gameServer := game.NewServer(gameMap)

	go gameServer.StartGameLoop()

	http.HandleFunc("/socket", socketHandler(gameServer))

	log.Printf("Start listening on: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
