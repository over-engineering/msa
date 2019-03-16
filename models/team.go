package models

// Team TBD
// We have to consider how to efficiently reflect user's
// own experience with globally fixed one. For example,
// the squad of teams, their winning history etc.
type Team interface {
	UID() UID
	Squad() []NPC
	Captain() NPC
	SelectCaptain() // 게임마다 스쿼드가 변화가 생기면 새로 그 경기의 주장을 뽑아서 임명해야 한다.
	Strategy() Strategy
	Stadium() Facility
	BestSquad() []NPC // 전략 및 스쿼드 구성원들 정보를 통해 베스트 스쿼드를 짜는 로직.
	Manager() NPC
	Finance()
	StarIndex() float32 // StartIndex range: 0~5
	History() []History
}

type Strategy interface {
}

type SoccerStrategy struct {
}

type SoccerTeam struct {
	ID       UID `json:"id"`
	squad    []NPC
	manager  NPC
	captain  NPC
	stadium  Facility
	strategy SoccerStrategy
	// TBD
}
