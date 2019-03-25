package models

type EmotionType int

const (
	Fun             EmotionType = iota // 재미
	Relief                             // 안심
	Satisfaction                       // 만족
	Dissatisfaction                    // 불만족
	Pleasure                           // 기쁨
	Sadness                            // 슬픔
	Rage                               // 분노
	Fear                               // 공포
)
