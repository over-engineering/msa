package space

import "github.com/over-engineering/msa/models/types"

type WeatherType types.Type

const (
	Sunny WeatherType = iota // 0
	Rainy
	Cloudy
)

// 습도, 미세먼지 농도 ...
