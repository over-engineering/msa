package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/over-engineering/msa/football"
	"github.com/over-engineering/msa/football/types"
)

func getAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ability, _ := football.FindAbilityByID(types.UID(params["id"]))
	json.NewEncoder(w).Encode(ability)
}

func createAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ability football.Ability
	_ = json.NewDecoder(r.Body).Decode(&ability)
	football.RegisterAbility(&ability)
	json.NewEncoder(w).Encode(ability)
}

func updateAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ability football.Ability
	_ = json.NewDecoder(r.Body).Decode(&ability)
	football.UpdateAbility(&ability)
	json.NewEncoder(w).Encode(ability)
}

func deleteAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	football.DeleteAbility(types.UID(params["id"]))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello World!")
// }

func main() {
	r := mux.NewRouter()

	// r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/api/ability/{id}", getAbility).Methods("GET")
	r.HandleFunc("/api/ability", createAbility).Methods("POST")
	r.HandleFunc("/api/ability", updateAbility).Methods("PUT")
	r.HandleFunc("/api/ability/{id}", deleteAbility).Methods("DELETE")

	// http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
