package system

import (
	common "temp/structure/common"
)

type RespiratorySystem struct {
	Nasal    common.Organ `json:"nasal"`
	Bronchus common.Organ `json:"bronchus"`
	Trachea  common.Organ `json:"trachea"`
	Lung     Lung         `json:"lung"`
}

type Lung struct {
	common.Organ
}
