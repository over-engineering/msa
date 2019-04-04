package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/over-engineering/msa/models/ability"
	"github.com/over-engineering/msa/models/starter"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := starter.CreateCharacter()
	fmt.Println(c)

	params := mux.Vars(r)
	switch params["game_type"] {
	case "football":
		ability := ability.UpdateFootballAbility(c.ID, &c.Status)

		req, _ := FootballClient.NewRequest("POST", "/api/ability", ability)
		fmt.Println(req, ability)
		// if err != nil {
		// 	return nil, err
		// }

		FootballClient.Do(req, &ability)
		json.NewEncoder(w).Encode(ability)
	default:
	}
}
