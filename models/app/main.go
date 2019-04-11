package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/over-engineering/msa/models/api"
)

const (
	FootballScheme = "http"
	FootballHost   = "localhost:8081"
)

func main() {
	api.FootballClient = api.NewClient(
		&url.URL{
			Scheme: FootballScheme,
			Host:   FootballHost,
		},
		"",
		nil,
	)

	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

	// 	ts := oauth2.StaticTokenSource(
	// 		&oauth2.Token{AccessToken: "... your access token ..."},
	// 	)
	// 	tc := oauth2.NewClient(oauth2.NoContext, ts)
	// 	client := github.NewClient(tc)

}
