package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/over-engineering/msa/common/models/training"
	"github.com/over-engineering/msa/common/models/types"
	"github.com/over-engineering/msa/common/models/user"
	"github.com/over-engineering/msa/common/starter"
)

type NewCharacterInfo struct {
	BornTime   string  `json:"born_time"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Generation int     `json:"generation"`
	Height     float32 `json:"height"`
	Weight     float32 `json:"weight"`
	Fat        float32 `json:"fat"`
}

func createCharacter(r *http.Request) (*user.Character, error) {
	var ci NewCharacterInfo
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	ncinfo, ok := r.PostForm["new_character_info"]
	if !ok {
		return nil, errors.New("new character info parameter expected")
	}
	if err := json.Unmarshal([]byte(ncinfo[0]), &ci); err != nil {
		return nil, err
	}
	c, err := starter.CreateNewCharacter(
		ci.BornTime,
		ci.FirstName,
		ci.LastName,
		ci.Generation,
		70, 40, 70, 40, 50, 30, 50, 30,
		ci.Height,
		ci.Weight,
		ci.Fat,
	)
	fmt.Println(c)
	if err != nil {
		return nil, err
	}
	return c, nil
	// json.NewEncoder(w).Encode(ability)
}

func apiPostCharacter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c, err := createCharacter(r)
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func UpdateCharacterByTraining(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tr := training.FindTrainingByID(types.UID(params["id"]))
	// c := character.FindCharacterByID(types.UID(params["id"]))
	tr.TakeTraining()
	// UpdateAbility(c.ID, params["game_type"], c.Status)
	// json.NewEncoder(w).Encode(ability)
}
