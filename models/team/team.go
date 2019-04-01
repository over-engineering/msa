package team

import (
	"errors"

	"github.com/over-engineering/msa/models/history"
	"github.com/over-engineering/msa/models/types"
)

// Team TBD
// We have to consider how to efficiently reflect user's
// own experience with globally fixed one. For example,
// the squad of teams, their winning history etc.
type Team struct {
	ID       types.UID         `json:"id"`
	Name     string            `json:"name"`
	Balance  float32           `json:"balance"` // Finance 관련 된 interface로 바꿀 수 있을 듯.
	History  []history.History `json:"history"`
	StarRate float32           `json:"star_rate"`
	// Nation
}

func makeNewTeam(id types.UID, name string, balance, starRate float32, history []history.History) *Team {
	return &Team{
		ID:       id,
		Name:     name,
		Balance:  balance,
		StarRate: starRate,
		History:  history,
	}
}

// Pay pays the amount of money from its balance
func (t *Team) Pay(amount float32) error {
	if t.Balance < amount {
		return errors.New("not enough balance for this entity")
	}
	t.Balance -= amount
	return nil
}

// Paid increases the amount of money in team's balance.
func (t *Team) Paid(amount float32) error {
	t.Balance += amount
	return nil
}

// GetName returns Team's name.
func (t Team) GetName() string {
	return t.Name
}

// GetID returns Team's ID.
func (t Team) GetID() types.UID {
	return t.ID
}

// FootballTeam represents the football team in the game.
type FootballTeam struct {
	Team
	// PlayerList []*FootballPlayer `json:"player_list"`
	// Manager    FootballManager   `json:"manager"`
	// Captain    *FootballPlayer   `json:"captain"`
	// Stadium    Stadium           `json:"stadium"`
	// TBD
}

// MakeNewFootballTeam returns new football team.
func MakeNewFootballTeam(
	id types.UID, name string, balance, starRate float32, history []history.History,
) *FootballTeam {
	team := makeNewTeam(id, name, balance, starRate, history)
	return &FootballTeam{
		Team: *team,
	}
}
