package physical

type ForceType int

const (
	Grip ForceType = iota // 0
	Kick
	Accelerating
)
