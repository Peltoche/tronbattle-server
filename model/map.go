package model

const MapSize = 1000

// -------------------------------------
// |             |-15tiles-|           |
// | |-15tiles-| user tile |-15 iles-| |
// |             |-15 iles-|		   |
// -------------------------------------
const ScreenSize = 31

type ElementKind int32

const (
	Empty ElementKind = iota
	Wall
	User
)

type Element struct {
	Kind  ElementKind       `json:"kind"`
	Metas map[string]string `json:"metas,omitempty"`
}
