package tests

import (
	"fmt"
	"net/url"
	"testing"

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

func TestUnity(t *testing.T) {
	client := api.NewClient(
		&url.URL{
			Scheme: CommonScheme,
			Host:   CommonHost,
		},
		"",
		nil,
	)

	req, err := client.NewRequest("POST", "/api/character/progamer", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req)
	client.Do(req, nil)
}
