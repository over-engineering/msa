package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/over-engineering/msa/models/api"
)

const (
	CommonScheme   = "http"
	CommonHost     = "localhost:8080"
	FootballScheme = "http"
	FootballHost   = "localhost:8081"
	ProgamerScheme = "http"
	ProgamerHost   = "localhost:8082"
)

func main() {
	// api.FootballClient = api.NewClient(
	// 	&url.URL{
	// 		Scheme: FootballScheme,
	// 		Host:   FootballHost,
	// 	},
	// 	"",
	// 	nil,
	// )
	api.ProgamerClient = api.NewClient(
		&url.URL{
			Scheme: ProgamerScheme,
			Host:   ProgamerHost,
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
