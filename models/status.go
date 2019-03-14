package models

// Status represents all aspects of character's current state.
type Status struct {
	// 신체 상태
	Health  float32               `json:"health"` // range: 0~100
	Height  float32               `json:"height"`
	Weight  float32               `json:"weight"`
	Joints  map[JointType]Joint   `json:"joints"`
	Muscles map[MuscleType]Muscle `json:"muscles"`
	Fat     map[FatType]Fat       `json:"fat"`

	// 수치적으로 계산이 들어갈 때 필요한 파라미터들
	Resilience float32 `json:"resilience"`
	Cardio     float32 `json:"cardio"`
	Charm      float32 `json:"charm"`

	Balance     float32 `json:"balance"`
	Decision    float32 `json:"decision"` // 판단력
	Agility     float32 `json:"agility"`
	Flexibility float32 `json:"flexibility"`
	// TODO: Sense

	// Mental related features
	Ambition      float32 `json:"ambition"`
	Boldness      float32 `json:"boldness"`
	Aggression    float32 `json:"aggression"`
	Predictation  float32 `json:"predictation"`
	Composure     float32 `json:"composure"` // 참착성
	Concentration float32 `json:"concentration"`
	Immersion     float32 `json:"immersion"`
	Competition   float32 `json:"competition"`
	SelfEsteem    float32 `json:"selfEsteem"`
	Confidence    float32 `json:"confidence"`
	Attention     float32 `json:"attention"` // 관종도

	// Intelligence(Gardner의 이론 차용)
	// 성향 이론
	// 성향은 나쁘고 좋은 게 없다. => 성향은 유저의 선택으로 영향을 받고 캐릭터의 행동을 결정 짓는 것 뿐.
	// 이벤트 관련 주변 연락, 약속 잡히는 빈도, 친구들의 성향.
	Intelligence map[IntelligenceType]Intelligence `json:"Intelligence"`
	Personality  map[PersonalityType]Personality   `json:"personality"`

	Memory     float32 `json:"memory"`
	Creativity float32 `json:"creativity"`
}
