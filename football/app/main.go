package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/over-engineering/msa/football"
	"github.com/over-engineering/msa/football/api"
)

func main() {
	es, _ := football.GetEquipmentsFromJSON()
	football.RegisterEquipments(es)
	for _, val := range football.EqMap {
		fmt.Println(*val)
	}

	ts, _ := football.GetTrainingsFromJSON()
	football.RegisterTrainings(ts)
	for _, val := range football.TrMap {
		fmt.Println(*val)
	}

	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
