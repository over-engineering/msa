package football

import "github.com/over-engineering/msa/football/types"

type TrainingMap map[types.UID]*Training

var TrMap = TrainingMap{}

type Trainings struct {
	Trainings []Training `json:"trainings"`
}

type Training struct {
	ID            types.UID               `json:"id"`
	TargetAbility map[AbilityType]float32 `json:"target_ability"`
	// Level         int                     `json:"level"`
	// Time          float32                 `json:"time"`
}

func RegisterTrainings(ts []Training) {
	for i, e := range ts {
		TrMap[e.ID] = &ts[i]
	}

	// TODO: training to db server
}

// func NewTraining(id types.UID, targetAbility map[AbilityType]float32) *Training {
// 	tr := &Training{ID: id, TargetAbility: targetAbility}
// 	TrMap[id] = tr
// 	return tr
// }

// 훈련 레벨, 시간, 컨디션에 따른 효율(0~1)에 따라 abiilty 상승이 결정됨
func (t *Training) TakeTraining(a *Ability, level int, time float32, efficiency float32) {
	vMap := map[AbilityType]float32{}
	for key, val := range t.TargetAbility {
		vMap[key] = efficiency * time * (val + float32(level))
	}
	a.AddValue(vMap)
}
