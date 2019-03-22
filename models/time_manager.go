package models

import (
	"fmt"
	"time"
)

type TimeManager struct {
	A int
}

func (t *TimeManager) CurrentTime() {
	fmt.Println(int32(time.Now().Unix()))
}
