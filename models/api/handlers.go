package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/over-engineering/msa/models/character"
	"github.com/over-engineering/msa/models/starter"
	"github.com/over-engineering/msa/models/training"
	"github.com/over-engineering/msa/models/types"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := starter.CreateCharacter()
	fmt.Println(c)
	params := mux.Vars(r)
	UpdateAbility(c.ID, params["game_type"], c.Status)
	// json.NewEncoder(w).Encode(ability)
}

func UpdateCharacterByTraining(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tr := training.FindTrainingByID(types.UID(params["id"]))
	c := character.FindCharacterByID(types.UID(params["id"]))
	tr.TakeTraining()
	UpdateAbility(c.ID, params["game_type"], c.Status)
	// json.NewEncoder(w).Encode(ability)
}
