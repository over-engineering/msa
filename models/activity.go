package models

import "fmt"

type ActivityManager interface {
	VisitFacility(f FacilityManager, options []int)
	ConsumeGoods(g GoodsHelper, options []int)
	PlayJob(j JobManager, options []int)
	HangOut(id UID)
}

const (
	MaxActPoint     = 100
	VisitActPoint   = -60
	ConsumeActPoint = -30
	PlayActPoint    = -30
	HangOutActPoint = -60
)

func ActVisit(a ActivityManager, f FacilityManager, options []int) error {
	// fmt.Println(options)
	err := AddActPoint(a, VisitActPoint)
	if err != nil {
		return err
	}
	a.VisitFacility(f, options)
	return nil
}

func ActConsume(a ActivityManager, g GoodsHelper, options []int) error {
	err := AddActPoint(a, ConsumeActPoint)
	if err != nil {
		return err
	}
	a.ConsumeGoods(g, options)
	return nil
}

func ActPlay(a ActivityManager, j JobManager, options []int) error {
	err := AddActPoint(a, PlayActPoint)
	if err != nil {
		return err
	}
	a.PlayJob(j, options)
	return nil
}

func ActHangOut(a ActivityManager, id UID) error {
	err := AddActPoint(a, HangOutActPoint)
	if err != nil {
		return err
	}
	a.HangOut(id)
	return nil
}

func AddActPoint(a ActivityManager, amount int) error {
	switch a.(type) {
	case *Character:
		if a.(*Character).ActPoint+amount < 0 {
			return fmt.Errorf("ActPoint < amount")
		}

		a.(*Character).ActPoint += amount
		if a.(*Character).ActPoint > MaxActPoint {
			a.(*Character).ActPoint = MaxActPoint
		}

	// case *NPC:
	default:
	}

	return nil
}
