package contract

import (
	"errors"
	"fmt"
	"time"

	"github.com/over-engineering/msa/models/finance"
	"github.com/over-engineering/msa/models/types"
)

// Contract interface represents all kinds of contracts in the game.
// 컨트랙트라는 건 어떤 entity와 entity 간의 약속. 즉 가장 기본적인 구조는 아래와
// 같이 갑, 을 + 해당 계약의 ID가 있다고 판단했음.
// 우선 항상 덩치로 생각할 때 A < B 라고 생각. A는 B와 계약을 맺는다. 이렇게 생각함.
type Contract struct {
	ID      types.UID `json:"id"`
	AID     types.UID `json:"aid"`
	BID     types.UID `json:"bid"`
	ASig    bool      `json:"asig"`
	BSig    bool      `json:"bsig"` // imageUrl을 입력해도 재밌을 듯
	Term    int       `json:"term"`
	ExpDate time.Time `json:"exp_date"`
}

func makeNewContract(aid, bid types.UID, term int) Contract {
	return Contract{
		AID:  aid,
		BID:  bid,
		ASig: false,
		BSig: false,
		Term: term,
	}
}

// Sign returns error during signing contract process.
func (c *Contract) Sign(id types.UID) error {
	switch id {
	case c.AID:
		c.ASig = true
		return nil
	case c.BID:
		c.BSig = true
		return nil
	default:
		return errors.New("id is not matching with any entity")
	}
}

// GetID returns contract's unique id.
func (c *Contract) GetID() types.UID {
	return c.ID
}

// SoccerPlayerContract represents the contract between soccer team
// and soccer player.
type SoccerPlayerContract struct {
	Contract   `json:"contract"`
	PlayerName string `json:"player_name"` // A
	TeamName   string `json:"team_name"`   // B

	Salary      finance.Dollars `json:"salary"`
	DownPayment finance.Dollars `json:"down_payment"` // 계약금
	Buyout      finance.Dollars `json:"buyout"`
	// Bonus Clauses
	GoalBonus   finance.Dollars `json:"goal_bonus"`
	AssistBonus finance.Dollars `json:"assist_bonus"`
	EntryBonus  finance.Dollars `json:"entry_bonus"`
	// League Win Bonus
	// Intercontinental Cup Win Bonus

	// Release type 논의 필요
	Violations []Violation `json:"violations"`
}

// MakeNewSoccerPlayerContract return new SoccerPlayerContract with
// given infos.
func MakeNewSoccerPlayerContract(
	a, b finance.Payer,
	term int,
	salary, downPayment finance.Dollars) *SoccerPlayerContract {
	contract := makeNewContract(a.GetID(), b.GetID(), term)
	return &SoccerPlayerContract{
		Contract:    contract,
		PlayerName:  a.GetName(),
		TeamName:    b.GetName(),
		Salary:      salary,
		DownPayment: downPayment,
	}
}

// Implement executes payment according to the contract.
func (sc *SoccerPlayerContract) Implement(b finance.Banker) error {
	b.TransferById(sc.Salary, sc.BID, sc.AID)
	return nil
}

// AddViolation adds new violation to contract
func (sc *SoccerPlayerContract) AddViolation(aToB bool, content ViolationType, fine float32) {
	switch aToB {
	case true: // Player = violator, Team = compensator
		newViolation := makeNewViolation(
			sc.Contract.AID,
			sc.Contract.BID,
			sc.PlayerName,
			sc.TeamName,
			content,
			fine)
		sc.Violations = append(sc.Violations, newViolation)
	default:
		newViolation := makeNewViolation(
			sc.Contract.BID,
			sc.Contract.AID,
			sc.TeamName,
			sc.PlayerName,
			content,
			fine)
		sc.Violations = append(sc.Violations, newViolation)
	}
}

// Describe returns the description of contract, this function
// makes SoccerPlayerContract to have Describer interfaceß
func (sc SoccerPlayerContract) Describe() string {
	// 우선은 한글로만. + 싸인 된 상황 고려하는 로직 필요.
	// playername과 teamname을 만든 이유는 선수 이름을 바꾸거나 팀이름이 바뀌
	// 더라도 계약서는 바뀌지 않아야하기 때문!
	description := sc.PlayerName + " 선수는 "
	description += sc.TeamName + " 팀과 "
	term := fmt.Sprintf("%d", sc.Term)
	salary := fmt.Sprintf("%.2f", sc.Salary)
	description += term + "시즌 동안 "
	description += "주급 " + salary + " 달러를 받기로 계약했습니다."
	// 달러랑 유로랑 원 등 계산 해야함.
	return description
}

// MonthlyPaymentContract represents all kinds of monthly paying contracts.
type MonthlyPaymentContract struct {
	Contract        `json:"contract"`
	UserName        string      `json:"user_name"`
	ServiceProvider string      `json:"service_provider"`
	Payment         float32     `json:"payment"`
	Date            int         `json:"date"`   // 2월 28일, 매달 말일은 따로 로직 필요.
	Unpaid          float32     `json:"unpaid"` // 밀린 돈 이자 계산 로직 나중에 추가.(신용도도)
	Violations      []Violation `json:"violations"`
	Availability    bool        `json:"availability"`
	ServiceType     ServiceType `json:"service_type"`
}

// ServiceType represents the specific type of services.
type ServiceType string

const (
	// PhoneService ...
	PhoneService ServiceType = "휴대폰 서비스"
	// InternetService ...
	InternetService ServiceType = "인터넷 서비스"
)

// MakeNewMonthlyPaymentContract returns new MonthlyPaymentContract.
func MakeNewMonthlyPaymentContract(
	aid, bid types.UID,
	uName, spName string,
	term, date int,
	payment float32,
	serviceType ServiceType,
) *MonthlyPaymentContract {
	contract := makeNewContract(aid, bid, term)
	return &MonthlyPaymentContract{
		Contract:        contract,
		UserName:        uName,
		ServiceProvider: spName,
		Payment:         payment,
		Date:            date,
		Availability:    true,
		ServiceType:     serviceType,
	}
}

// ExecPayment executes payment according to the contract.
func (mc *MonthlyPaymentContract) ExecPayment() error {
	// 1. Find character or team by ID
	// 2. Substitute the amount of payment from paying account
	//    If account's balance is not enough, increase Unpaid value
	//    and make availability to false(3 out 제도 생각해보기.)
	//    If not, add payment value to paid account.(Tax cal. required)
	return nil
}

// Describe returns the description of service.
func (mc *MonthlyPaymentContract) Describe() string {
	term := fmt.Sprintf("%d", mc.Term)
	date := fmt.Sprintf("%d", mc.Date)
	payment := fmt.Sprintf("%.2f", mc.Payment)
	description := mc.UserName + "은(는) "
	description += mc.ServiceProvider + " 로부터 "
	description += string(mc.ServiceType) + "를 "
	description += term + "년 동안 "
	description += "매달 " + date + "일에 " + payment + " 달러를 내고 사용한다."

	return description
}

// Violation represents all kinds of violations with contract.
type Violation struct {
	ViolatorID    types.UID     `json:"aid"`
	CompensatorID types.UID     `json:"bid"`
	Violator      string        `json:"violator"`
	Compensator   string        `json:"compensator"`
	Content       ViolationType `json:"content"`
	Fine          float32       `json:"fine"`
}

// ViolationType represents the specific type of violation.
type ViolationType string

const (
	// TrainingAbsence ...
	TrainingAbsence ViolationType = "훈련 불참"
	// SalaryUnpaid ...
	SalaryUnpaid ViolationType = "주급 미지급"
)

func makeNewViolation(
	vID, cID types.UID,
	violator, compensator string,
	content ViolationType,
	fine float32) Violation {
	return Violation{
		ViolatorID:    vID,
		CompensatorID: cID,
		Violator:      violator,
		Compensator:   compensator,
		Content:       content,
		Fine:          fine,
	}
}

// Describe returns the description of violation.
func (v Violation) Describe() string {
	fine := fmt.Sprintf("%.2f", v.Fine)
	description := v.Violator + "은(는) " // 요거 간단하게 할 수 있는 로직 있음.
	description += string(v.Content) + " 시 "
	description += v.Compensator + "에게 "
	description += fine + "달러의 벌금을 내야한다."
	return description
}
