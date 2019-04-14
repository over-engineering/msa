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
		"CreateCharacter",
		"POST",
		"/api/character/{game_type}",
		apiPostCharacter,
	},
	Route{
		"UpdateCharacterByTraining",
		"PUT",
		"/api/character/training/{id}/{game_type}",
		UpdateCharacterByTraining,
	},
}
