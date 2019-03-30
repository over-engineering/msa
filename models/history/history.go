package history

import "time"

// History represents the interface for all kinds of histories
type History interface {
}

type LeagueHistory struct {
}

type TeamHistory struct {
	League []LeagueHistory `json:"league_won"`
	// Rival       *Team           `json:"rival"`
	// StarPlayers []*Entity       `json:"star_players"`
	Created time.Time `json:"created"`
}
