package models

// type Effect struct {
// 	Target string      `json:"target"` // test 요망
// 	Value  interface{} `json:"value"`
// }

type Effect struct {
	Target     string        `json:"target"` // test 요망
	Value      interface{}   `json:"value"`
	Parameters []interface{} `json:"parameters"`
}

type Effects []Effect

func (es Effects) ApplyEffects() {
	for _, e := range es {
		switch {
		case e.Target == "UpdateIntelligences":
			e.Value.(func(m map[IntelligenceType]float32))(e.Parameters[0].(map[IntelligenceType]float32))
		case e.Target == "UpdatePersonality":
			e.Value.(func(ps Personality))(e.Parameters[0].(Personality))
		case e.Target == "UpdateMentals":
			e.Value.(func(values Mentals, p Personality))(e.Parameters[0].(Mentals), e.Parameters[1].(Personality))
		// case e.Target == "Intelligence":
		// 	s.Intelligence.UpdateValue(Value.(Intelligence))

		case e.Target == "AddAllFriendships":
			// fmt.Println(e.Value, e.Parameters)
			e.Value.(func(obj1 interface{}, value float32))(e.Parameters[0].(*Character), e.Parameters[1].(float32))
		default:
			// f := reflect.ValueOf(s).Elem().FieldByName(e.Target)
			// f.SetFloat(f.Float() + e.Value.(float64))
		}
	}

}
