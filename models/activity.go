package models

import "fmt"

type ActivityManager interface {
	VisitFacility(f FacilityManager, options []int)
	ConsumeGoods(g GoodsHelper, options []int)
	PlayJob(j JobManager, options []int)
	HangOut(id UID)
}

const (
	VisitActPoint   = -60
	ConsumeActPoint = -30
	PlayActPoint    = -30
	HangOutActPoint = -60
)

func ActVisit(a ActivityManager, f FacilityManager, options []int) error {
	a.VisitFacility(f, options)
	return AddActPoint(a, VisitActPoint)
}

func ActConsume(a ActivityManager, g GoodsHelper, options []int) error {
	a.ConsumeGoods(g, options)
	return AddActPoint(a, ConsumeActPoint)
}

func ActPlayform(a ActivityManager, j JobManager, options []int) error {
	a.PlayJob(j, options)
	return AddActPoint(a, PlayActPoint)
}

func ActHangOut(a ActivityManager, id UID) error {
	a.HangOut(id)
	return AddActPoint(a, HangOutActPoint)
}

func AddActPoint(a ActivityManager, amount int) error {
	switch a.(type) {
	case *Character:
		if a.(*Character).ActPoint < amount {
			return fmt.Errorf("a.ActivityPoint < amount")
		}
		a.(*Character).ActPoint += amount

	// case *NPC:
	default:
	}

	return nil
}
