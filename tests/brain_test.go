package tests

import (
	"fmt"
	"testing"

	"github.com/over-engineering/msa/models"
)

func TestBrain(t *testing.T) {
	st := models.StarterStatus()
	fmt.Println("======================Brain test=========================")
	fmt.Println(st)
	fmt.Println("Update brain")
	es := models.Effects{
		models.Effect{
			Target: "Brain",
			Value: models.MathBrain{
				models.MathmaticalCalculation: 10,
				models.MathmaticalApprox:      10,
			},
		},
	}
	st.ApplyEffects(es)
	fmt.Println(st)
	fmt.Println("=========================================================")

}
