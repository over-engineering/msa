package starter

import (
	"math/rand"
	"time"

	"github.com/over-engineering/msa/common/models/fame"
	"github.com/over-engineering/msa/common/models/finance"

	"github.com/over-engineering/msa/common/models/entity"
	"github.com/over-engineering/msa/common/models/job"
	"github.com/over-engineering/msa/common/models/status"
	"github.com/over-engineering/msa/common/models/talent"
	"github.com/over-engineering/msa/common/models/types"
	"github.com/over-engineering/msa/common/models/user"
)

const initialValue = float32(50)

func CreateNewCharacter(
	bornTime string,
	fName, lName string,
	generation int,
	pBase, pSigma, bBase, bSigma, nBase, nSigma, hBase, hSigma float32,
	height, weight, fat float32,
) (*user.Character, error) {
	born, err := time.Parse(time.RFC3339, bornTime)
	if err != nil {
		return nil, err
	}
	current, _ := time.Parse(time.RFC3339, "2013-01-02T15:04:05+09:00")
	s := status.NewStatus(
		born,
		current,
		generation,
		pBase, pSigma,
		bBase, bSigma,
		nBase, nSigma,
		hBase, hSigma,
		height, weight, fat,
	)
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	t := talent.NewTalents(r, 100, 100)
	i, err := entity.NewIndividual(
		fName,
		lName,
		job.None,
		[]status.Updater{},
		[]types.UID{},
		[]types.UID{},
		[]types.UID{},
		*s,
		t,
		fame.FanInfo{
			Fan:  2,
			Anti: 0,
		},
		finance.Finance{
			Balance: 1000,
		},
	)
	if err != nil {
		return nil, err
	}
	c := user.NewCharacter(*i, current)
	return c, nil
}

// var StatusType1 = character.Status{
// 	lifecycle.LifeCycle `json:"life_cycle"` // range: 0~
// 	// Physical
// 	physical.Physical `json:"physical"`
// 	// Brain
// 	brain.Brain `json:"brain"`
// 	// Needs
// 	needs.Needs    `json:"needs"`
// 	needs.Hormones `json:"hormones"`

// 	LifeCycle: LifeCycle{Age: 20, Dead: time.Now().AddDate(10, 5, 5)},
// 	Health:    100,
// 	Height:    170,
// 	Weight:    70,
// 	// Joints
// 	// Muscles
// 	// Fat
// 	// Resilience
// 	// Cardio
// 	// Charm
// 	// Balance
// 	// Decision
// 	// Agility
// 	// Flexibility
// 	// Mentals: character.Mentals{
// 	// 	Ambition:     &Mental{Type: Ambition, Value: 50},
// 	// 	Boldness:     &Mental{Type: Boldness, Value: 50},
// 	// 	Aggression:   &Mental{Type: Aggression, Value: 50},
// 	// 	Predictation: &Mental{Type: Predictation, Value: 50},
// 	// 	Composure:    &Mental{Type: Composure, Value: 50},
// 	// 	// Concentration: &Mental{Type: Concentration, Value: 50},
// 	// 	Immersion:   &Mental{Type: Immersion, Value: 50},
// 	// 	Competition: &Mental{Type: Competition, Value: 50},
// 	// 	SelfEsteem:  &Mental{Type: SelfEsteem, Value: 50},
// 	// 	Confidence:  &Mental{Type: Confidence, Value: 50},
// 	// 	Attention:   &Mental{Type: Attention, Value: 50},
// 	// },
// 	// Intelligences: Intelligences{
// 	// 	MusicalRhythmic:     &Intelligence{Type: MusicalRhythmic, Value: 50},
// 	// 	VisualSpatial:       &Intelligence{Type: VisualSpatial, Value: 50},
// 	// 	VerbalLinguistic:    &Intelligence{Type: VerbalLinguistic, Value: 50},
// 	// 	LogicalMathematical: &Intelligence{Type: LogicalMathematical, Value: 50},
// 	// 	BodilyKinesthetic:   &Intelligence{Type: BodilyKinesthetic, Value: 50},
// 	// 	Interpersonal:       &Intelligence{Type: Interpersonal, Value: 50},
// 	// 	Intrapersonal:       &Intelligence{Type: Intrapersonal, Value: 50},
// 	// 	Naturalistic:        &Intelligence{Type: Naturalistic, Value: 50},
// 	// },
// 	// Personality: NewPersonality([4]float32{0.5, 0.4, 0.6, 0.8}),
// 	// Memory
// 	// Creativity
// }

// const Num = 100
// const MaxRange = float32(10)
// const MinRange = float32(-10)

// func StarterStatus() *character.Status {

// 	var randList = [Num]float32{}
// 	for i := 0; i < Num; i++ {
// 		randList[i] = rand.Float32()*(MaxRange-MinRange) + MinRange
// 	}

// 	NewStatus := &character.Status{

// 		Talents: NewTalents([]float32{randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50}),
// 		// LifeCycle: LifeCycle{Age: randList[0] + 20, LifeSpan: randList[1] + 100, Generation: 0},
// 		Brain: Brain{
// 			AthleticBrain: NewSubBrain(0, []float32{randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50}),
// 			MusicBrain:    NewSubBrain(1, []float32{randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50}),
// 			LanguageBrain: NewSubBrain(2, []float32{randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50}),
// 			MathBrain:     NewSubBrain(3, []float32{randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50}),
// 			VisualBrain:   NewSubBrain(4, []float32{randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50}),
// 		},
// 		// Talents: NewTalents([]float32{randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50, randList[0] + 50}),
// 		// LifeCycle: LifeCycle{Age: randList[0] + 20, LifeSpan: randList[1] + 100, Generation: 0},

// 		Health: randList[2] + 50,
// 		Height: randList[3] + 170,
// 		Weight: randList[4] + 70,
// 		// Joints
// 		// Muscles: 50
// 		// Fat: 50
// 		KcKgTranslationRate: 0.8 + randList[24]/100,
// 		ConsumedKcal:        0,
// 		// Resilience:          randList[4] + 50,
// 		// Cardio:              randList[4] + 50,
// 		// Charm:               randList[4] + 50,
// 		// Balance:             randList[4] + 50,
// 		// Decision:            randList[4] + 50,
// 		// Agility:             randList[4] + 50,
// 		// Flexibility:         randList[4] + 50,
// 		Needs: Needs{
// 			FoodNeeds:        0,
// 			SleepNeeds:       0,
// 			SexNeeds:         0,
// 			RestNeeds:        0,
// 			StabilityNeeds:   0,
// 			AffectionNeeds:   0,
// 			ReputationNeeds:  0,
// 			AchievementNeeds: 0,
// 		},
// 		Hormones: Hormones{
// 			// Serotonin:      new(float32),
// 			// Norepinephrine: new(float32),
// 			// Epinephrine:    new(float32),
// 			// Dopamine:       new(float32),
// 			// Endorphin:      new(float32),
// 			// Oxytocin:       new(float32),
// 			Serotonin:      float32(0),
// 			Norepinephrine: float32(0),
// 			Epinephrine:    float32(0),
// 			Dopamine:       float32(0),
// 			Endorphin:      float32(0),
// 			Oxytocin:       float32(0),
// 		},
// 		Mentals: Mentals{
// 			Ambition:     &Mental{Type: Ambition, Value: randList[5] + 50},
// 			Boldness:     &Mental{Type: Boldness, Value: randList[6] + 50},
// 			Aggression:   &Mental{Type: Aggression, Value: randList[7] + 50},
// 			Predictation: &Mental{Type: Predictation, Value: randList[8] + 50},
// 			Composure:    &Mental{Type: Composure, Value: randList[9] + 50},
// 			// Concentration: &Mental{Type: Concentration, Value: randList[10] + 50},
// 			Immersion:   &Mental{Type: Immersion, Value: randList[11] + 50},
// 			Competition: &Mental{Type: Competition, Value: randList[12] + 50},
// 			SelfEsteem:  &Mental{Type: SelfEsteem, Value: randList[13] + 50},
// 			Confidence:  &Mental{Type: Confidence, Value: randList[14] + 50},
// 			Attention:   &Mental{Type: Attention, Value: randList[15] + 50},
// 		},
// 		// Intelligences: Intelligences{
// 		// 	MusicalRhythmic:     &Intelligence{Type: MusicalRhythmic, Value: randList[16] + 50},
// 		// 	VisualSpatial:       &Intelligence{Type: VisualSpatial, Value: randList[17] + 50},
// 		// 	VerbalLinguistic:    &Intelligence{Type: VerbalLinguistic, Value: randList[18] + 50},
// 		// 	LogicalMathematical: &Intelligence{Type: LogicalMathematical, Value: randList[19] + 50},
// 		// 	BodilyKinesthetic:   &Intelligence{Type: BodilyKinesthetic, Value: randList[20] + 50},
// 		// 	Interpersonal:       &Intelligence{Type: Interpersonal, Value: randList[21] + 50},
// 		// 	Intrapersonal:       &Intelligence{Type: Intrapersonal, Value: randList[22] + 50},
// 		// 	Naturalistic:        &Intelligence{Type: Naturalistic, Value: randList[23] + 50},
// 		// },
// 		Personality: NewPersonality([4]float32{randList[24]/100 + 0.5, randList[25]/100 + 0.4, randList[26]/100 + 0.6, randList[27]/100 + 0.8}),
// 		// Memory
// 		// Creativity
// 	}
// 	return NewStatus
// }
