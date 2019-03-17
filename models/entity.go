package models

// Entity represents all kinds of entities in the game.
type Entity interface {
	GetID() UID
	FullName() string
	GetStatus() // status could be interface
	GetEffect()
	GetContract() Contract
}
