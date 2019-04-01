package models

// FatType is an enum for various fat types
type FatType Type

// Fat represents the fats in character's body
type Fat struct {
	Mass float32 `json:"mass"` // Mass represents the mass of fat
}
