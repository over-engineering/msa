package system

import (
	common "temp/structure/common"
)

type CirculatorySystem struct {
	Heart       Heart       `json:"heart"`
	BloodVessel BloodVessel `json:"blood_vessel`
}

type Heart struct {
	common.Organ

	Rate float32 `json:"rate"` // Rate represents the heart rate
}

type BloodVessel struct {
	common.Organ
}
