package models

import (
	"reflect"
)

// Status represents all aspects of character's current state.
type Status struct {
	// 재능
	Talents Talents `json:"talents"`

	// 나이
	LifeCycle LifeCycle `json:"life_cycle"` // range: 0~

	// brain
	Brain []BrainManager `json:"brain"`

	// 신체 상태
	Health  float32               `json:"health"` // range: 0~100
	Height  float32               `json:"height"`
	Weight  float32               `json:"weight"`
	Joints  map[JointType]Joint   `json:"joints"`
	Muscles map[MuscleType]Muscle `json:"muscles"`
	Fat     map[FatType]Fat       `json:"fat"`

	// 위치 옮겨야 함
	KcKgTranslationRate float32 `json:"kc_kg_translate_rate"`
	ConsumedKcal        float32

	// 수치적으로 계산이 들어갈 때 필요한 파라미터들
	// Resilience float32 `json:"resilience"`
	// Cardio     float32 `json:"cardio"`
	// Charm      float32 `json:"charm"`

	// Balance     float32 `json:"balance"`
	// Decision    float32 `json:"decision"` // 판단력
	// Agility     float32 `json:"agility"`
	// Flexibility float32 `json:"flexibility"`
	// TODO: Sense

	// Mental related features
	Needs    Needs    `json:"needs"`
	Hormones Hormones `json:"hormones"`
	Mentals  Mentals  `json:"mentals"`

	// Intelligence(Gardner의 이론 차용)
	// 성향 이론
	// 성향은 나쁘고 좋은 게 없다. => 성향은 유저의 선택으로 영향을 받고 캐릭터의 행동을 결정 짓는 것 뿐.
	// 이벤트 관련 주변 연락, 약속 잡히는 빈도, 친구들의 성향.
	// Intelligences Intelligences `json:"Intelligence"`
	Personality Personality `json:"personality"`
}

func (s *Status) ApplyEffects(es Effects) {
	for _, e := range es {
		switch {
		case e.Target == "Talents":
			s.Talents.UpdateValue(e.Value.(Talents))
		case e.Target == "Brain":
			UpdateBrain(s.Brain, e.Value.(BrainManager), s.Talents)
		case e.Target == "Personality":
			s.Personality.UpdateValue(e.Value.(Personality))
		case e.Target == "Needs":
			s.Needs.UpdateValue(e.Value.(Needs), s.Hormones)
		case e.Target == "Hormones":
			s.Hormones.UpdateValue(e.Value.(Hormones))
		case e.Target == "Mentals":
			s.Mentals.UpdateValue(e.Value.(Mentals), s.Personality)
		default:

			// f := reflect.ValueOf(s).Elem().FieldByName(e.Target)
			// f.SetFloat(f.Float() + e.Value.(float64))
			v := reflect.ValueOf(s).Elem().FieldByName(e.Target)
			u := v.Interface().(float32)
			v.SetFloat(float64(u + e.Value.(float32)))

		}
	}
}

func (s *Status) UpdateByDay() {
	// Update needs
	nMap := Needs{
		FoodNeeds:  FoodNeedsRisingVal,
		SleepNeeds: SleepNeedsRisingVal,
		SexNeeds:   SexNeedsRisingVal,
	}
	s.Needs.UpdateValue(nMap, s.Hormones)

	// Update hormones
	// s.Hormones.UpdateValue()

	// Update height

	// Update weight
	s.Weight += s.ConsumedKcal * s.KcKgTranslationRate
	s.ConsumedKcal = 0
}
