package db

import "math/big"

type Quest struct {
	QuestId    int         `json:"quest_id"`
	Type       int         `json:"type"`
	IsAdvTask  int         `json:"is_adv_task"`
	Next       []int       `json:"next"`
	Condition  []*big.Int  `json:"condition"`
	QuestAward map[int]int `json:"questAward"`
}

type QuestData struct {
	QuestId  int32
	Status   int32
	Complete []*big.Int
}
