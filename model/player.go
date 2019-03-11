package model

import "github.com/gorilla/websocket"

type Direction int32

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Player struct {
	Username  string
	Direction Direction
	Position  Position
	Conn      *websocket.Conn
}
