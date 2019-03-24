package models

import (
	"strings"
)

// Character represents the game character belongs to user
// Things could be referencing with unique IDs.
type Character struct {
	ID        int64   `json:"id,string"`
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
	Goods    [][]GoodsHelper `json:"goods"`
	Finance  Finance         `json:"finance"` // Or maybe Account
	Contract Contract        `json:"contract"`
	FanInfo  FanInfo         `json:"fan_info"`

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

// FullName returns character's full name. This function have to
// consider user's linguistic setting.
func (c *Character) FullName() string {
	// TODO: user's  linguistic setting check
	return strings.Join([]string{c.FirstName, c.LastName}, " ")
}

func (c *Character) VisitFacility(f FacilityManager, options []int) {
	f.Visit(c.Location, c.Status)

	switch f.(type) {
	case *Stadium:

	case *Market:
		if options[0] == 0 {
			g := f.(*Market).BuyGoods(options[1])
			RegisterGoods(c.Goods, g)
		} else if options[0] == 1 {
			g := GetGoods(c.Goods, GoodsType(options[2]), options[3])
			f.(*Market).SellGoods(g)
			UnRegisterGoods(c.Goods, GoodsType(options[2]), options[3])
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
	case *Food:
		if options[0] == 0 {
			g.(*Food).Eat(c.Status)
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
