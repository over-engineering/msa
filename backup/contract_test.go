package models

import "fmt"

func ExampleMakeNewSoccerPlayerContract() {
	a := &Entity{
		ID:      "Player#001",
		Name:    "박찬",
		Balance: float32(1000.0),
	}
	b := &Entity{
		ID:      "Team#001",
		Name:    "바르샤",
		Balance: float32(100000000.0),
	}
	term := 2
	salary := float32(10000.0)
	downPayment := float32(1200.0)

	newContract := MakeNewSoccerPlayerContract(
		a, b, term, salary, downPayment,
	)
	fmt.Println(newContract.Contract)
	newContract.Sign(a.ID)
	newContract.Sign(b.ID)
	fmt.Println(newContract.Contract)
	newContract.AddViolation(true, TrainingAbsence, float32(200))
	newContract.AddViolation(false, SalaryUnpaid, float32(15000.0))
	PrintDescription(newContract)
	newContract.Violations[0].Fine -= 14.0 // 협상 로직에 넣어야 할 것.
	for _, v := range newContract.Violations {
		PrintDescription(v)
	}
	// Output:
	// { Player#001 Team#001 false false 2 0001-01-01 00:00:00 +0000 UTC}
	// { Player#001 Team#001 true true 2 0001-01-01 00:00:00 +0000 UTC}
	// 박찬 선수는 바르샤 팀과 2시즌 동안 주급 10000.00 달러를 받기로 계약했습니다.
	// 박찬은(는) 훈련 불참 시 바르샤에게 186.00달러의 벌금을 내야한다.
	// 바르샤은(는) 주급 미지급 시 박찬에게 15000.00달러의 벌금을 내야한다.
}

func ExampleMakeNewMonthlyPaymentContract() {
	aid := UID("Player#001")
	bid := UID("PhoneSP#001")
	uName := "박찬"
	spName := "SKT"
	term := 2
	date := 15
	payment := float32(50)
	serviceType := PhoneService

	newMonthlyPaymentContract := MakeNewMonthlyPaymentContract(
		aid, bid, uName, spName, term, date, payment, serviceType,
	)
	PrintDescription(newMonthlyPaymentContract)
	// Output:
	// 박찬은(는) SKT 로부터 휴대폰 서비스를 2년 동안 매달 15일에 50.00 달러를 내고 사용한다.
}

func ExampleMakeNewFootballTeam() {
	id := UID("FT#001")
	name := "바르샤"
	balance := float32(1000000000)
	starRate := float32(5.0)
	history := []History{}

	newTeam := MakeNewFootballTeam(id, name, balance, starRate, history)
	a := &Entity{
		ID:      "Player#001",
		Name:    "박찬",
		Balance: float32(1000.0),
	}
	term := 2
	salary := float32(10000.0)
	downPayment := float32(1200.0)

	newContract := MakeNewSoccerPlayerContract(
		a, newTeam, term, salary, downPayment,
	)
	fmt.Println(newContract.Contract)
	newContract.Sign(a.ID)
	newContract.Sign(newTeam.ID)
	newTeam.PlayerList = append(newTeam.PlayerList, a)
	fmt.Println(newContract.Contract)
	newContract.AddViolation(true, TrainingAbsence, float32(200))
	newContract.AddViolation(false, SalaryUnpaid, float32(15000.0))
	PrintDescription(newContract)
	newContract.Violations[0].Fine -= 14.0 // 협상 로직에 넣어야 할 것.
	for _, v := range newContract.Violations {
		PrintDescription(v)
	}
	newTeam.Pay(newContract.Salary)
	a.Paid(newContract.Salary)
	fmt.Println(newTeam.Balance)
	fmt.Println(*newTeam.PlayerList[0])
	fmt.Println(*a)

	// Output:
	// { Player#001 FT#001 false false 2 0001-01-01 00:00:00 +0000 UTC}
	// { Player#001 FT#001 true true 2 0001-01-01 00:00:00 +0000 UTC}
	// 박찬 선수는 바르샤 팀과 2시즌 동안 주급 10000.00 달러를 받기로 계약했습니다.
	// 박찬은(는) 훈련 불참 시 바르샤에게 186.00달러의 벌금을 내야한다.
	// 바르샤은(는) 주급 미지급 시 박찬에게 15000.00달러의 벌금을 내야한다.
	// 9.9999e+08
	// {Player#001 박찬 11000}
	// {Player#001 박찬 11000}
}
