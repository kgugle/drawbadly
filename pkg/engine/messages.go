package engine

import ()

// PixelMessage ...
type PixelMessage struct {
	X           uint32
	Y           uint32
	Color       uint8
	StrokeStart bool
}

type PlayerAction int

const (
	JOINED = iota
	LEFT
)

// PlayerUpdateMessage  ...
type PlayerUpdateMessage struct {
	PlayerID int
	Action   PlayerAction
}

// PlayerConnectionMessage  ...
type PlayerConnectionMessage struct {
	PlayerID            int
	NewConnectionStatus int
}

// ScoreUpdateMessage
type ScoreUpdateMessage struct {
	PlayerID int
	NewScore int
}
