package socket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/tronbattle/server/game"
	"github.com/tronbattle/server/model"
)

type Connection struct {
	conn       *websocket.Conn
	gameServer *game.Server
	id         string
}

func NewConnection(conn *websocket.Conn, gameServer *game.Server) *Connection {
	return &Connection{
		conn:       conn,
		gameServer: gameServer,
		id:         uuid.NewV4().String(),
	}
}

func (t *Connection) ListenInputEvents() {
	for {
		var cmd model.Command

		err := t.conn.ReadJSON(&cmd)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Print("client deconnected")
			} else {
				log.Printf("failed to read a message: %s", err)
			}

			break
		}

		switch cmd.Command {
		case "start":
			err = t.dispatchStartCommand(cmd.Body)
		default:
			fmt.Printf("unknown command: %s -> %s\n", cmd.Command, cmd.Body)
		}

		if err != nil {
			log.Printf("failed to execute the command: %s", err)
		}
	}
}

func (t *Connection) dispatchStartCommand(rawCmd json.RawMessage) error {
	var cmd model.StartCmd

	err := json.Unmarshal(rawCmd, &cmd)
	if err != nil {
		return errors.Wrap(err, "failed to parse the command body for the \"start\" cmd")
	}

	t.gameServer.EnterGame(t.id, t.conn, &cmd)

	return nil
}
