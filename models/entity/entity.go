package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/over-engineering/msa/models/fame"
	"github.com/over-engineering/msa/models/finance"
	"github.com/over-engineering/msa/models/job"
	"github.com/over-engineering/msa/models/status"
	"github.com/over-engineering/msa/models/talent"
	"github.com/over-engineering/msa/models/types"
	world "github.com/over-engineering/msa/models/world/space"
)

type Individual struct {
	ID           types.UID        `json:"id,string"`
	FirstName    string           `json:"first_name"`
	LastName     string           `json:"last_name"`
	JobType      job.Type         `json:"job_type"`
	Location     world.Location   `json:"location"`
	Conditions   []status.Updater `json:"conditions"`
	Contracts    []types.UID      `json:"contracts"`
	GoodsList    []types.UID      `json:"goods_list"`
	FacilityList []types.UID      `json:"facility_list"`

	status.Status   `json:"status"`
	talent.Talents  `json:"talents"`
	fame.FanInfo    `json:"fan_info"`
	finance.Finance `json:"finance"`
}

// AddCondition adds updater interface to Conditions
func (i *Individual) AddCondition(u status.Updater) error {
	i.Conditions = append(i.Conditions, u)
	u.ApplyEffect(&i.Status)
	return nil
}

// RemoveCondition removes updater from conditions
func (i *Individual) RemoveCondition(index int) error {
	if index >= len(i.Conditions) {
		return errors.New("index out of range")
	}
	i.Conditions[index].CancelEffect(&i.Status)
	i.Conditions = append(i.Conditions[:index], i.Conditions[index+1:]...)
	return nil
}

// UpdateCondition updates conditions with given time
func (i *Individual) UpdateCondition(ct time.Time) error {
	for index, condition := range i.Conditions {
		if !condition.Validate(ct) {
			err := i.RemoveCondition(index)
			return err
		}
	}
	return nil
}

// CurrentJob returns individual's current job.
// From this result we could adjust UI, events eti.
func (i Individual) CurrentJob() string {
	return i.JobType.String()
}

// GetID returns Entity's ID.
func (i Individual) GetID() types.UID {
	return i.ID
}

// GetName returns character's full name. This function have to
// consider user's linguistic setting.
func (i Individual) GetName() string {
	// TODO: user's  linguistic setting check
	return strings.Join([]string{i.FirstName, i.LastName}, " ")
}

// Pay pays the amount of money from its balance
func (i *Individual) Pay(amount finance.Dollars) error {
	return i.SubBalance(amount)
}

// Paid increases the amount of money in entity's balance.
func (i *Individual) Paid(amount finance.Dollars) error {
	i.AddBalance(amount)
	return nil
}

// Corporation represents 주식회사.
type Corporation struct {
	ID   types.UID `json:"id,string"`
	Name string    `json:"name"`
	// 지분 정보
	fame.FanInfo    `json:"fan_info"`
	finance.Finance `json:"finance"`

	FacilityList []types.UID `json:"facility_list"`
}

func (c Corporation) GetID() types.UID {
	return c.ID
}

func (c Corporation) GetName() string {
	return c.Name
}

func (c *Corporation) Pay(amount finance.Dollars) error {
	return c.SubBalance(amount)
}

// Paid increases the amount of money in entity's balance.
func (c *Corporation) Paid(amount finance.Dollars) error {
	c.AddBalance(amount)
	return nil
}

type Pet struct {
	ID   types.UID `json:"id,string"`
	Name string    `json:"name"`
}
