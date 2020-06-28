package engine

import (
	"time"
)

// MessageType ...
type MessageType int

const (
	PixelMessageType MessageType = iota
	ChatMessageType
	PlayerConnMessageType
	CorrectGuessMessageType
	ScoreMessageType
	RoundUpdateMessageType
)

type SocketMessage interface {
	GetMessageType() MessageType
	// TODO: marshal/unmarshal methods for bytes instead of using json
}

// PixelMessage ...
type PixelMessage struct {
	X           uint32
	Y           uint32
	Color       uint8
	StrokeStart bool
}

func (m *PixelMessage) GetMessageType() MessageType {
	return PixelMessageType
}

// ChatMessage ...
type ChatMessage struct {
	Payload  string
	SentTime time.Time
}

func (m *ChatMessage) GetMessageType() MessageType {
	return ChatMessageType
}

// PlayerConnectionMessage  ...
type PlayerConnectionMessage struct {
	PlayerID            int
	NewConnectionStatus ConnectionStatus
}

func (m *PlayerConnectionMessage) GetMessageType() MessageType {
	return PlayerConnMessageType
}

// CorrectGuessMessage ...
type CorrectGuessMessage struct {
	PlayerID int
	Received int
}

func (m *CorrectGuessMessage) GetMessageType() MessageType {
	return CorrectGuessMessageType
}

// ScoreUpdateMessage also marks end of round
type ScoreUpdateMessage struct {
	NewScores []int
}

func (m *ScoreUpdateMessage) GetMessageType() MessageType {
	return ScoreMessageType
}

// RoundUpdateMessage ...
type RoundUpdateMessage struct {
	SecondsLeft int
	NewDrawerId int // only populated at start of round
}

func (m *RoundUpdateMessage) GetMessageType() MessageType {
	return RoundUpdateMessageType
}
