package models

import "fmt"

type ActivityManager struct {
	Character *Character
	Facility  Facility
	Job       interface{}
}

func (a *ActivityManager) Visit(f Facility) {
	// switch f.(type) {
	// case Stadium:
	// default:
	// }
	dis := Distance(f.GetLocation(), a.Character.CurrentLocation())
	a.Character.Status.ApplyEffects(
		Effects{
			Effect{
				Target: "Health",
				Value:  -30 * dis,
			},
		})
}

func (a *ActivityManager) HosService(h *Hospital, option int) {
	switch option {
	case 1:
		h.Treat(a.Character.Status, a.Character.Conditions)
	case 2:
		h.Surgery(a.Character.Status, a.Character.Conditions)
	default:
	}
}

func (a *ActivityManager) FitService(f *Fitness, option int) {
	switch option {
	case 1:
	default:
	}
}

func (a *ActivityManager) BuyGoods(g GoodsIt) {
	fmt.Println("Buy goods")
	g.RegisterGoods(a.Character)
}

func (a *ActivityManager) SellGoods(g GoodsIt) {
	g.UnRegisterGoods(a.Character)
}

func (a *ActivityManager) Call(p *SmartPhone, option int) {
	switch option {
	case 1:
		p.Call(a.Character.Status)
	default:
	}
}

func (a *ActivityManager) Ride(c *Car, option int) {
	switch option {
	case 1:
		c.Ride(a.Character.Status)
	default:
	}
}

func (a *ActivityManager) Eat(f *Food, option int) {
	switch option {
	case 1:
		f.Eat(a.Character.Status)
	default:
	}

}

func (a *ActivityManager) HangOut(id UID) {

}

func (a *ActivityManager) Rest() {
	a.Character.Status.ApplyEffects(
		Effects{
			Effect{
				Target: "Health",
				Value:  30,
			},
		})
}
