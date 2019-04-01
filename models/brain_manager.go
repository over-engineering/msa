package models

type BrainManager interface {
	UpdateValue(vMap interface{}, tal Talents)
	UpdateByAge(age int)
}

func UpdateBrain(bl []BrainManager, b interface{}, t Talents) {
	switch b.(type) {
	case AthleticBrain:
		bl[Athletic].UpdateValue(b, t)
	case MusicBrain:
		bl[Music].UpdateValue(b, t)
	case LanguageBrain:
		bl[Language].UpdateValue(b, t)
	case MathBrain:
		bl[Math].UpdateValue(b, t)
	case VisualBrain:
		bl[Visualization].UpdateValue(b, t)
	default:
	}
}
