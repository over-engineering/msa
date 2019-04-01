package physical

import "github.com/over-engineering/msa/models/types"

// FatType is an enum for various fat types
type FatType types.Type

// Fat represents the fats in character's body
type Fat struct {
	Mass float32 `json:"mass"` // Mass represents the mass of fat
}
