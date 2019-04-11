package models

// import "github.com/over-engineering/msa/models/types"

// // type Friendship float32
// type Friendship struct {
// 	// PersonalityType PersonalityType `json:"peronalityType"`
// 	Value float32 `json:"value"`
// }
// type Friendships map[types.UID]*Friendship

// // 성격에 따른 계수로 같은 같은 성격이면 최대 시너지,
// // 정반대의 성격이면 최소값을 가짐
// var FriendshipTable = [PersonalityNum][PersonalityNum]float32{
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// 	[PersonalityNum]float32{1.1, 1, 0.9, 1.2, 1.3, 0.8, 0.7, 1.5, 0.5, 1, 1, 1.1, 1.2, 0.9, 0.8, 1},
// }

// // func NewFriendship(sourcePs PersonalityType, targetPs PersonalityType, value float32) *Friendship {
// // 	f := Friendship(FriendshipTable[sourcePs-1][targetPs-1] * value)
// // 	return &f
// // }

// // func (f *Friendship) AddFriendship(sourcePs PersonalityType, targetPs PersonalityType, value float32) {
// // 	*f += Friendship(FriendshipTable[sourcePs-1][targetPs-1] * value)
// // }

// // func NewFriendship(sourcePs PersonalityType, targetPs PersonalityType, value float32) *Friendship {
// // 	f := Friendship{
// // 		PersonalityType: targetPs,
// // 		Value:           FriendshipTable[sourcePs-1][targetPs-1] * value,
// // 	}
// // 	return &f
// // }

// // func (f *Friendship) AddFriendship(sourcePs PersonalityType, value float32) {
// // 	(*f).Value += FriendshipTable[sourcePs-1][(*f).PersonalityType-1] * value
// // }

// // func AddFriendships(obj1 interface{}, obj2 interface{}, value float32) {
// // 	// func AddFriendships(obj1 interface{}, id UID, value float32) {
// // 	var id1 EntityID
// // 	// id2 := id
// // 	var id2 EntityID
// // 	var ps1 PersonalityType
// // 	var ps2 PersonalityType
// // 	var fs1 Friendships
// // 	var fs2 Friendships

// // 	switch v := obj1.(type) {
// // 	case *Character:
// // 		id1 = (v.ID)
// // 		ps1 = v.Status.Personality.Type
// // 		fs1 = v.Friendships
// // 	case *Entity:
// // 	default:

// // 	}

// // 	// switch v := obj2.(type) {
// // 	// case *Character:
// // 	// 	id2 = UID(v.ID)
// // 	// 	ps2 = v.Status.Personality.Type
// // 	// 	fs2 = v.Friendships
// // 	// case *Entity:
// // 	// default:

// // 	// }

// // 	if _, ok := fs1[id2]; ok {
// // 		fs1[id2].AddFriendship(ps1, value)
// // 	} else {
// // 		fs1[id2] = NewFriendship(ps1, ps2, value)
// // 	}

// // 	if _, ok := fs2[id1]; ok {
// // 		fs2[id1].AddFriendship(ps1, value)
// // 	} else {
// // 		fs2[id1] = NewFriendship(ps1, ps2, value)
// // 	}
// // }

// func AddAllFriendships(obj1 interface{}, value float32) {
// 	var ps1 PersonalityType
// 	var fs1 Friendships

// 	switch v := obj1.(type) {
// 	case *Character:
// 		ps1 = v.Status.Personality.Type
// 		fs1 = v.Friendships
// 	case *Entity:
// 	default:

// 	}

// 	for key, _ := range fs1 {
// 		fs1[key].AddFriendship(ps1, value)
// 	}
// }
