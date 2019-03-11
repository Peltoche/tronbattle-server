package game

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tronbattle/server/model"
)

type Server struct {
	players map[string]*model.Player
	gameMap *Map
}

func NewServer(gameMap *Map) *Server {
	return &Server{
		players: map[string]*model.Player{},
		gameMap: gameMap,
	}
}

func (t *Server) StartGameLoop() {
	for {
		for id := range t.players {
			t.movePlayerForward(t.players[id])
		}

		for id := range t.players {
			screenMap := t.retrieveScreenMapForPlayer(id)

			t.sendScreen(t.players[id].Conn, screenMap)
		}

		time.Sleep(time.Second)
	}
}

func (t *Server) EnterGame(userID string, conn *websocket.Conn, cmd *model.StartCmd) {
	player := model.Player{
		Username:  cmd.Username,
		Direction: model.Down,
		//Position:  t.gameMap.SelectRandomStartPoint(),
		Position: model.Position{X: 10, Y: 10},
		Conn:     conn,
	}

	t.players[userID] = &player
	fmt.Printf("new player: %v\n", player)
}

func (t *Server) UpdateDirection(userID string, direction model.Direction) {
	t.players[userID].Direction = direction
}

func (t *Server) movePlayerForward(player *model.Player) {

	switch player.Direction {
	case model.Down:
		player.Position.Y += 1
	case model.Up:
		player.Position.Y -= 1
	case model.Right:
		player.Position.X += 1
	case model.Left:
		player.Position.X -= 1
	default:
		log.Printf("unknown player direction %#v", player.Direction)
	}
	log.Printf("player pos: %+v", player.Position)
}

func (t *Server) retrieveScreenMapForPlayer(userID string) [model.ScreenSize * model.ScreenSize]model.Element {
	res := [model.ScreenSize * model.ScreenSize]model.Element{}

	player := t.players[userID]
	pos := player.Position

	minX := pos.X - (model.ScreenSize / 2)
	maxX := pos.X + (model.ScreenSize / 2)

	minY := pos.Y - (model.ScreenSize / 2)
	maxY := pos.Y + (model.ScreenSize / 2)

	log.Printf("min/max: %v/%v - %v/%v", minX, maxX, minY, maxY)
	screenX := 0
	screenY := 0
	for x := minX; x < maxX; x++ {
		if x >= 0 && x <= model.MapSize {
			for y := minY; y < maxY; y++ {
				if y >= 0 && y <= model.MapSize {
					res[screenX*model.ScreenSize+screenY] = t.gameMap.inner[x*model.MapSize+y]
				}

				screenY += 1
			}
		}

		screenY = 0
		screenX += 1
	}

	return res
}

func (t *Server) sendScreen(conn *websocket.Conn, screenMap [model.ScreenSize * model.ScreenSize]model.Element) {
	rawScreenMap, err := json.Marshal(screenMap)
	if err != nil {
		log.Printf("failed marshal the map for map update: %s", err)
		return
	}

	err = conn.WriteJSON(&model.Event{
		EventType: "mapUpdate",
		Body:      rawScreenMap,
	})
	if err != nil {
		log.Printf("failed to write a map update: %s", err)
	}
}
