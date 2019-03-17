package models

import "time"

// History represents the interface for all kinds of histories
type History interface {
}

type LeagueHistory struct {
}

type TeamHistory struct {
	League      []LeagueHistory `json:"league_won"`
	Rival       *Team           `json:"rival"`
	StarPlayers []*NPC          `json:"star_players"`
	Created     time.Time       `json:"created"` // Maybe we have to make our own timestamp

}
