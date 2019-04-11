package models

type DayManager interface {
	UpdateByDay()
}

func UpdateDay(ds ...DayManager) {
	for _, d := range ds {
		d.UpdateByDay()
	}
}
