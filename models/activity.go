package models

import "fmt"

type ActivityManager struct {
	Character *Character
	JobHelper JobHelper
	ActPoint  int // 0 ~ 100
}

const MinActivityPoint = 30

func (a *ActivityManager) Act(amount int) error {
	if a.ActPoint < amount {
		return fmt.Errorf("a.ActivityPoint < amount")
	}

	a.ActPoint -= amount
	return nil
}

func (a *ActivityManager) Visit(f Facility) error {
	if e := a.Act(50); e != nil {
		return e
	}

	dis := Distance(f.GetLocation(), a.Character.CurrentLocation())
	a.Character.Status.ApplyEffects(
		Effects{
			Effect{
				Target: "Health",
				Value:  -30 * dis,
			},
		})
	return nil
}

func (a *ActivityManager) MarketService(m *Market, option int) error {
	if e := a.Act(10); e != nil {
		return e
	}

	switch option {
	case 1:
		// ShowGoodsList()
	case 2:
		sel := 0
		g := m.BuyGoods(sel)
		RegisterGoods(a.Character.Goods, g)
	case 3:
		goodsType := FoodType
		goodsIndex := 0
		g := GetGoods(a.Character.Goods, goodsType, goodsIndex)
		m.SellGoods(g)
		UnRegisterGoods(a.Character.Goods, goodsType, goodsIndex)
	default:
	}

	return nil
}

func (a *ActivityManager) HosService(h *Hospital, option int) error {
	if e := a.Act(50); e != nil {
		return e
	}

	switch option {
	case 1:
		h.Treat(a.Character.Status, a.Character.Conditions)
	case 2:
		h.Surgery(a.Character.Status, a.Character.Conditions)
	default:
	}

	return nil
}

func (a *ActivityManager) FitService(f *Fitness, option int) error {
	if e := a.Act(50); e != nil {
		return e
	}

	switch option {
	case 1:
		f.Exercise(a.Character.Status)
	default:
	}

	return nil
}

// func (a *ActivityManager) BuyGoods(g GoodsHelper) error {
// 	if e := a.Act(50); e != nil {
// 		return e
// 	}

// 	// fmt.Println("Buy goods")
// 	g.RegisterGoods(a.Character)
// 	return nil
// }

// func (a *ActivityManager) SellGoods(g GoodsHelper) error {
// 	if e := a.Act(50); e != nil {
// 		return e
// 	}

// 	g.UnRegisterGoods(a.Character)
// 	return nil
// }

func (a *ActivityManager) Call(p *SmartPhone, option int) error {
	if e := a.Act(50); e != nil {
		return e
	}

	switch option {
	case 1:
		p.Call(a.Character.Status)
	default:
	}

	return nil
}

func (a *ActivityManager) Ride(c *Car, option int) error {
	if e := a.Act(50); e != nil {
		return e
	}

	switch option {
	case 1:
		c.Ride(a.Character.Status)
	default:
	}
	return nil
}

func (a *ActivityManager) Eat(f *Food, option int) error {
	if e := a.Act(50); e != nil {
		return e
	}

	switch option {
	case 1:
		f.Eat(a.Character.Status)
	default:
	}

	return nil
}

func (a *ActivityManager) HangOut(id UID) error {
	if e := a.Act(50); e != nil {
		return e
	}

	a.Character.Status.ApplyEffects(
		Effects{
			Effect{
				Target: "Charm",
				Value:  5,
			},
		})

	// c2 := GetCharacterByUID(UID) // read from db?
	// AddFriendships(a.Character, c2, 10)

	return nil
}

func (a *ActivityManager) Rest() error {
	if e := a.Act(50); e != nil {
		return e
	}

	a.Character.Status.ApplyEffects(
		Effects{
			Effect{
				Target: "Health",
				Value:  30,
			},
		})

	return nil
}

func (a *ActivityManager) Training() error {
	if e := a.Act(50); e != nil {
		return e
	}

	a.JobHelper.Training()
	return nil
}

func (a *ActivityManager) Play() error {
	if e := a.Act(50); e != nil {
		return e
	}

	a.JobHelper.Play()
	return nil
}
