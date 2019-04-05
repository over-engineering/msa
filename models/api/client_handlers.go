package api

import (
	"fmt"

	"github.com/over-engineering/msa/models/character"

	"github.com/over-engineering/msa/models/ability"
	"github.com/over-engineering/msa/models/types"
)

var FootballClient *Client

func UpdateAbility(id types.UID, gameType string, status character.Status) {
	switch gameType {
	case "football":
		ability := ability.UpdateFootballAbility(id, &status)
		req, _ := FootballClient.NewRequest("POST", "/api/ability", ability)
		fmt.Println(req, ability)
		// if err != nil {
		// 	return nil, err
		// }

		FootballClient.Do(req, &ability)
	default:
	}
}
