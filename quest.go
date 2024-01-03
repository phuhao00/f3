package main

import "frame_3/db"

type Condition interface {
	Quest() *Quest
	GetIdx() int
	CheckComplete(player *Player, data *db.QuestData) bool
	Complete(player *Player)
	Attach(player *Player, data *db.QuestData)
	Detach(player *Player)
}

type Quest struct {
	quest     *db.Quest
	Condition []Condition
}

func (q *Quest) Quest() *Quest {
	//TODO implement me
	panic("implement me")
}

func (q *Quest) GetIdx() int {
	//TODO implement me
	panic("implement me")
}

func (q *Quest) CheckComplete(player *Player, data *db.QuestData) bool {
	for _, v := range q.Condition {
		v.Attach(player, data)
	}

	flg := true
	for _, v := range q.Condition {
		if !v.CheckComplete(player, data) {
			flg = false
			break
		}
	}

	if flg {
		data.Status = 1
		//todo notify client
	} else if true /* sync*/ {
		//todo notify client
	}
	return true
}

func (q *Quest) Complete(player *Player) {

}

func (q *Quest) Attach(player *Player, data *db.QuestData) {
	q.Check(player, data, false)
}

func (q *Quest) Detach(player *Player) {

}

func (this *Quest) Check(player *Player, data *db.QuestData, sync bool) {

}
