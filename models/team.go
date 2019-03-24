package models

// Team TBD
// We have to consider how to efficiently reflect user's
// own experience with globally fixed one. For example,
// the squad of teams, their winning history etc.
type Team interface {
	UID() UID
	Squad() []Entity
	GetCaptain() Entity
	SelectCaptain() // 게임마다 스쿼드가 변화가 생기면 새로 그 경기의 주장을 뽑아서 임명해야 한다.
	GetStrategy() Strategy
	MainPlace() FacilityManager
	BestSquad() []Entity // 전략 및 스쿼드 구성원들 정보를 통해 베스트 스쿼드를 짜는 로직.
	GetManager() Entity
	Finance()
	StarIndex() float32 // StartIndex range: 0~5
	History() []History
}

type Strategy interface {
}

type SoccerStrategy struct {
}

type SoccerTeam struct {
	ID       UID            `json:"id"`
	Members  []Entity       `json:"members"`
	Manager  Entity         `json:"manager"`
	Captain  Entity         `json:"captain"`
	Stadium  Stadium        `json:"stadium"`
	Strategy SoccerStrategy `json:"strategy"`
	// TBD
}
