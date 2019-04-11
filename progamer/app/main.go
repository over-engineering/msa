package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/over-engineering/msa/progamer"
	"github.com/over-engineering/msa/progamer/api"
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

	es, _ := progamer.GetEquipmentsFromJSON()
	progamer.RegisterEquipments(es)
	for _, val := range progamer.EqMap {
		fmt.Println(*val)
	}

	ts, _ := progamer.GetTrainingsFromJSON()
	progamer.RegisterTrainings(ts)
	for _, val := range progamer.TrMap {
		fmt.Println(*val)
	}

	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8082", router))

	// 	ts := oauth2.StaticTokenSource(
	// 		&oauth2.Token{AccessToken: "... your access token ..."},
	// 	)
	// 	tc := oauth2.NewClient(oauth2.NoContext, ts)
	// 	client := github.NewClient(tc)

}
