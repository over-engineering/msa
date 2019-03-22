package tests

import (
	"testing"

	"github.com/over-engineering/msa/models"
)

func TestTime(t *testing.T) {
	time := &models.TimeManager{A: 1}
	time.CurrentTime()
}
