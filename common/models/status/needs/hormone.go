package needs

import (
	"fmt"
	"math/rand"

	"github.com/over-engineering/msa/common/models/types"
)

type HormoneType types.Type

const (
	Serotonin HormoneType = iota //  세로토닌, 마음의 평안(긍정적 생각), 세로토닌은 행복을 느끼는 데에 기여한다고 일반적으로 생각되고 있다.
	// This is the one you can blame for PMS and your moody teenager. Serotonin controls your mood, appetite, and your sleep cycles
	// Low level of serotonin makes you impulsive and aggressive. Serotonin helps you to remain calm
	Norepinephrine // Too much gives you an anxious feeling while too little can leave you feeling depressed or sedated.
	Epinephrine    // 아드레날린, 흥분 Ever felt worried in exams? Fear,Rage, Amusement.. Almost all extreme emotions are caused by these two hormones. You may have heard the term Adrenalin Rush. This hormone is responsible for it.
	Dopamine       // 도파민, 도박 당첨 됐을 때 쾌감
	Endorphin      // 엔돌핀, 기쁨
	Oxytocin       // 옥시토신, 최고의 행복감
)

// type Hormones map[HormoneType]float32
type Hormones []float32

func NewHormones(r *rand.Rand, b, s float32) Hormones {
	hormones := Hormones{}
	for i := Serotonin; i <= Oxytocin; i++ {
		hormones = append(hormones, (r.Float32()-0.5)*s+b)
	}
	return hormones
}

func (hs Hormones) UpdateValue(v Hormones) {
	for key, value := range v {
		hs[key] += value
		fmt.Println("Update Hormone", key, value, hs[key])
	}
}
