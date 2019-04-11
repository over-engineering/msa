package models

import (
	"fmt"
	"time"
)

func ExampleMakeNewCondition() {
	name := "감기"
	duration := time.Duration(10000000000000)
	conditionType1 := MakeNewConditionType(
		name, &duration, nil,
	)
	PrintDescription(conditionType1)
	newCondition := MakeNewCondition(
		*conditionType1, time.Now(),
	)
	PrintDescription(*newCondition)
	// fmt.Println(newCondition.endTime)
	newCondition.Extend(duration)
	// fmt.Println(newCondition.endTime)
	newCondition.Shorten(2 * duration)
	// fmt.Println(newCondition.endTime)

	newConditions := Conditions{
		name: *newCondition,
	}
	newConditions.DeleteByTime(time.Now())
	fmt.Println(newConditions)
	// Output:
	// 감기은(는) 보통 2h46m40s 동안 나타납니다.
	// 감기은(는) 보통 2h46m40s 동안 나타납니다.
	// map[]
}
