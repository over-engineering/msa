package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/over-engineering/msa/football"
	"github.com/over-engineering/msa/football/types"
)

// func CreateCharacter(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	football.CreateCharacter(CommonClient)
// }

func GetAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ability, _ := football.FindAbilityByID(types.UID(params["id"]))
	json.NewEncoder(w).Encode(ability)
}

func CreateAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ability football.Ability
	_ = json.NewDecoder(r.Body).Decode(&ability)
	football.RegisterAbility(&ability)
	football.RegisterEquipList(ability.ID)
	fmt.Println(ability)
	json.NewEncoder(w).Encode(ability)
}

func UpdateAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ability football.Ability
	_ = json.NewDecoder(r.Body).Decode(&ability)
	football.UpdateAbility(&ability)
	json.NewEncoder(w).Encode(ability)
}

func DeleteAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	football.UnRegisterAbility(types.UID(params["id"]))
	football.UnRegisterEquipList(types.UID(params["id"]))
}

func GetTraining(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	training, _ := football.FindTrainingByID(types.UID(params["id"]))
	json.NewEncoder(w).Encode(training)
}

func TakeTraining(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ability, _ := football.FindAbilityByID(types.UID(params["character_id"]))
	training, _ := football.FindTrainingByID(types.UID(params["training_id"]))
	level, _ := strconv.Atoi(params["level"])
	time, _ := strconv.ParseFloat(params["time"], 32)
	efficiency, _ := strconv.ParseFloat(params["efficiency"], 32)
	training.TakeTraining(ability, level, float32(time), float32(efficiency))
	json.NewEncoder(w).Encode(ability)
}

func GetEquipment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	equipment, _ := football.FindEquipmentByID(types.UID(params["id"]))
	json.NewEncoder(w).Encode(equipment)
}

func EquipEquipment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ability, _ := football.FindAbilityByID(types.UID(params["character_id"]))
	eList, _ := football.FindEquipListByCharacterID(types.UID(params["character_id"]))
	equipment, _ := football.FindEquipmentByID(types.UID(params["equipment_id"]))
	err := equipment.Equip(ability, eList)
	fmt.Println(ability, err)
	json.NewEncoder(w).Encode(ability)
}

func UnEquipEquipment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ability, _ := football.FindAbilityByID(types.UID(params["character_id"]))
	eList, _ := football.FindEquipListByCharacterID(types.UID(params["character_id"]))
	equipment, _ := football.FindEquipmentByID(types.UID(params["equipment_id"]))
	err := equipment.UnEquip(ability, eList)
	fmt.Println(ability, err)
	json.NewEncoder(w).Encode(ability)
}
