package api

import (
	"fmt"

	"github.com/over-engineering/msa/common/models/ability"
	"github.com/over-engineering/msa/common/models/status"
	"github.com/over-engineering/msa/common/models/types"
)

var FootballClient *Client
var ProgamerClient *Client

func UpdateAbility(id types.UID, gameType string, status status.Status) {
	switch gameType {
	case "football":
		ability := ability.UpdateFootballAbility(id, &status)
		req, _ := FootballClient.NewRequest("POST", "/api/ability", ability)
		fmt.Println(req, ability)
		// if err != nil {
		// 	return nil, err
		// }

		FootballClient.Do(req, &ability)
	case "progamer":
		ability := ability.UpdateProgamerAbility(id, &status)
		req, _ := ProgamerClient.NewRequest("POST", "/api/ability", ability)
		fmt.Println(req, ability)
		// if err != nil {
		// 	return nil, err
		// }

		ProgamerClient.Do(req, &ability)
	default:
	}
}
