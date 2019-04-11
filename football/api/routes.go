package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"GetAbility",
		"GET",
		"/api/ability/{id}",
		GetAbility,
	},
	Route{
		"CreateAbility",
		"POST",
		"/api/ability",
		CreateAbility,
	},
	Route{
		"UpdateAbility",
		"PUT",
		"/api/ability",
		UpdateAbility,
	},
	Route{
		"DeleteAbility",
		"DELETE",
		"/api/ability/{id}",
		DeleteAbility,
	},
	Route{
		"GetEquipment",
		"GET",
		"/api/equipment/{id}",
		GetEquipment,
	},
	Route{
		"EquipEquipment",
		"GET",
		"/api/equipment/{character_id}/{equipment_id}",
		EquipEquipment,
	},
	Route{
		"UnEquipEquipment",
		"GET",
		"/api/equipment/{character_id}/{equipment_id}",
		UnEquipEquipment,
	},
	Route{
		"GetTraining",
		"GET",
		"/api/training/{id}",
		GetTraining,
	},
	Route{
		"TakeTraining",
		"GET",
		"/api/training/{character_id}/{training_id}/{level}/{time}/{efficiency}",
		TakeTraining,
	},
}
