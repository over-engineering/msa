package football

import (
	"fmt"
	"testing"

	// "github.com/mitchellh/mapstructure"
	"github.com/over-engineering/msa/football"
)

func TestJson(t *testing.T) {
	// data := football.ReadJsonFile("equipment.json")
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

}
