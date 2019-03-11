package model

import "encoding/json"

type Event struct {
	EventType string          `json:"ev_type"`
	Body      json.RawMessage `json:"body"`
}

type UpdateMapEvent struct {
	Map [ScreenSize * ScreenSize]Element `json:"map"`
}
