package models

// Entity represents all kinds of entities in the game.
// type Entity struct {
// 	ID      UID     `json:"id"`
// 	Name    string  `json:"name"`
// 	Balance float32 `json:"balance"`
// }

// // Pay pays the amount of money from its balance
// func (e *Entity) Pay(amount float32) error {
// 	if e.Balance < amount {
// 		return errors.New("not enough balance for this entity")
// 	}
// 	e.Balance -= amount
// 	return nil
// }

// // Paid increases the amount of money in entity's balance.
// func (e *Entity) Paid(amount float32) error {
// 	e.Balance += amount
// 	return nil
// }

// // GetName returns Entity's name.
// func (e Entity) GetName() string {
// 	return e.Name
// }

// // GetID returns Entity's ID.
// func (e Entity) GetID() UID {
// 	return e.ID
// }

// // FootballManager represents manager for soccer.
// type FootballManager struct {
// 	Entity
// }

// // FootballPlayer represents football player.
// type FootballPlayer struct {
// 	Entity
// 	Contract *SoccerPlayerContract
// }

// // GetContract returns contract for FootballPlayer
// func (f FootballPlayer) GetContract() *SoccerPlayerContract {
// 	return f.Contract
// }

// // type Coach struct {
// // 	Entity
// // }

// // type Owner struct {
// // 	Entity
// // }

// // type Friend struct {
// // 	Entity
// // }
