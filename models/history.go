package models

import "time"

// History represents the interface for all kinds of histories
type History interface {
}

type LeagueHistory struct {
}

type TeamHistory struct {
	League      []LeagueHistory `json:"leagueWon"`
	Rival       *Team           `json:"rival"`
	StarPlayers []*NPC          `json:"starPlayers"`
	Created     time.Time       `json:"created"` // Maybe we have to make our own timestamp

}
