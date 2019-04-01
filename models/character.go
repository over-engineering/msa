package models

import (
	"strings"
)

// Character represents the game character belongs to user
// Things could be referencing with unique IDs.
type Character struct {
	ID        UID     `json:"id,string"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Job       JobType `json:"job_type"`
	// TODO: Nationality
	// TODO: user struct

	Location *Location `json:"location"`

	Status     *Status                     `json:"status"`
	Conditions Conditions                  `json:"conditions"`
	Equipments map[EquipmentType]Equipment `json:"equipments"`
	// TODO: 논의 Equipments map[EquipmentType]UUID

	// Goods []map[GoodsHelper]GoodsHelper `json:"goods"`
	// Goods    []map[interface{}]interface{} `json:"goods"`
	// Goods    [][]GoodsHelper `json:"goods"`
	// Goods    []map[UID]GoodsHelper `json:"goods"`
	Goods    GoodsList `json:"goods"`
	Finance  Finance   `json:"finance"` // Or maybe Account
	Contract Contract  `json:"contract"`
	FanInfo  FanInfo   `json:"fan_info"`

	Friendships Friendships `json:"friendships"`
	// Character would not have a team. Also, team information
	// could be shared with other players.
	Team     *Team `json:"team"`
	ActPoint int   `json:"act_point"`
}

// NewCharacter returns new Character object
func NewCharacter() *Character {
	return &Character{}
}

// CurrentJob returns character's current job.
// From this result we could adjust UI, events etc.
func (c *Character) CurrentJob() JobType {
	return c.Job
}

func (c *Character) CurrentLocation() *Location {
	return c.Location
}

// CurrentContract returns character's current contract.
func (c *Character) CurrentContract() Contract {
	return c.Contract
}

// CurrentTeam returns character's current team.
func (c *Character) CurrentTeam() *Team {
	return c.Team
}

// GetName returns character's full name. This function have to
// consider user's linguistic setting.
func (c Character) GetName() string {
	// TODO: user's  linguistic setting check
	return strings.Join([]string{c.FirstName, c.LastName}, " ")
}

// // Pay pays the amount of money from its balance
// func (c *Character) Pay(amount float32) error {
// 	if c.Balance < amount {
// 		return errors.New("not enough balance for this entity")
// 	}
// 	c.Balance -= amount
// 	return nil
// }

// // Paid increases the amount of money in entity's balance.
// func (c *Character) Paid(amount float32) error {
// 	c.Balance += amount
// 	return nil
// }

// GetID returns Entity's ID.
func (c Character) GetID() UID {
	return c.ID
}

func (c *Character) VisitFacility(f FacilityManager, options []int) {
	f.Visit(c.Location, c.Status)

	switch f.(type) {
	case *Stadium:

	case *Market:
		if options[0] == 0 {
			g := f.(*Market).BuyGoods(options[1])
			c.Goods.RegisterGoods(g)
		} else if options[0] == 1 {
			g := c.Goods.GetGoods(GoodsType(options[1]), options[2])
			f.(*Market).SellGoods(g)
			c.Goods.UnRegisterGoods(GoodsType(options[1]), options[2])
		}
	case *Hospital:
		if options[0] == 0 {
			f.(*Hospital).Treat(c.Status, c.Conditions)
		} else if options[0] == 1 {
			f.(*Hospital).Surgery(c.Status, c.Conditions)
		}
	case *Fitness:
		if options[0] == 0 {
			f.(*Fitness).Exercise(c.Status)
		}
	default:
	}
}

func (c *Character) ConsumeGoods(g GoodsHelper, options []int) {
	switch g.(type) {
	case *House:
		if options[0] == 0 {
			g.(*House).Rest(c.Status)
		}
	case FoodManager:
		if options[0] == 0 {
			g.(FoodManager).Eat(c.Status)
		}
	case *SmartPhone:
		if options[0] == 0 {
			g.(*SmartPhone).Call(c.Status)
		}
	case *Car:
		if options[0] == 0 {
			g.(*Car).Ride(c.Status)
		}
	default:
	}

	return
}

func (c *Character) PlayJob(j JobManager, options []int) {
	switch j.(type) {
	case *SoccerPlayer:
	default:
	}
}

func (c *Character) HangOut(id UID) {
	c.Status.ApplyEffects(
		Effects{
			Effect{
				Target: "Charm",
				Value:  5,
			},
		})

}
