package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/over-engineering/msa/football"
	"github.com/over-engineering/msa/football/api"
)

const (
	CommonScheme = "http"
	CommonHost   = "localhost:8080"
)

// var CommanClient api.Client

func main() {
	api.CommonClient = api.NewClient(
		&url.URL{
			Scheme: CommonScheme,
			Host:   CommonHost,
		},
		"",
		nil,
	)

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
	log.Fatal(http.ListenAndServe(":8081", router))

	// 	ts := oauth2.StaticTokenSource(
	// 		&oauth2.Token{AccessToken: "... your access token ..."},
	// 	)
	// 	tc := oauth2.NewClient(oauth2.NoContext, ts)
	// 	client := github.NewClient(tc)

	
}
