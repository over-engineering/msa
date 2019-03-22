package models

// JobType is an enum for various job types.
type JobType Type

const (
	JOB_NONE            JobType = iota // 0
	JOB_SOCCER_PLAYER                  // 1
	JOB_BASEBALL_PLAYER                // 2
)

type JobHelper interface {
	RegisterPro()
	UnRegisterPro()
	Training()
	Play()
}

type SoccerPlayer struct {
	Ability SoccerAbility
}

func (s *SoccerPlayer) RegisterPro() {

}

func (s *SoccerPlayer) UnRegisterPro() {

}

func (s *SoccerPlayer) Training() {

}

func (s *SoccerPlayer) Play() {

}

type SoccerAbility struct {
}
