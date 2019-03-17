package models

type Effect struct {
	Target string      `json:"target"` // test 요망
	Value  interface{} `json:"value"`
}

type Effects []Effect
