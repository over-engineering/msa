package models

// Contract represents all kinds of contracts in the game.
type Contract interface {
	Sign()
	// Negotiate(c Contract) bool
}

type SoccerTeamContract struct {
	Team        Team      `json:"team"`
	Salary      float32   `json:"salary"`
	DownPayment float32   `json:"down_payment"` // 계약금
	Penalties   []Penalty `json:"penalties"`
	Options     []Option  `json:"options"`
}

// Penalty could be diverse from early contract termination to violation
// for team training.
type Penalty interface {
	Violation() string // 어떤 것을 위반했는지를 string으로 반환
	Value() int        // 벌금이 얼마인지를 반환
}

// Option could be diverse from simple incentives to release, buyout
// condition.
type Option interface {
	Type() Type // Option type을 반환
	// 여기다가 이제 option이 발동하는 컨디션 및 결과를 지정해주는 메소드를 지정하고자 함.
}
