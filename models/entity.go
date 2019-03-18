package models

// Entity represents all kinds of entities in the game.
type Entitier interface {
	GetID() UID
	FullName() string
	GetStatus() // status could be interface
	GetEffect()
	GetContract() Contract
}

// Basic struct for entity
type Entity struct {
	ID          UID         `json:"id"`
	Name        string      `json:"name"`
	Status      Status      `json:"status"`
	Description string      `json:"description"`
	Effects     Effects     `json:"effects"`
	Contract    *Contract   `json:"contract"`
	Friendships Friendships `json:"friendships`
}

//
type Manager struct {
	Entity
}

type Coach struct {
	Entity
}

type Owner struct {
	Entity
}

type Friend struct {
	Entity
}
