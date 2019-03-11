package model

import "encoding/json"

type Command struct {
	Command string          `json:"cmd"`
	Body    json.RawMessage `json:"body"`
}

type StartCmd struct {
	Username string `json:"username"`
}
