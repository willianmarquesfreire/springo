package domain

import (
	"time"
	"springo/core"
	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	GenericDomain `json:",omitempty" bson:",inline"`
	Information  string `json:"information"`
	ExpirateAt   time.Time `json:"expirateAt,omitempty"`
	User         *User `json:"user,omitempty"`
	Group     *Group `json:"group,omitempty"`
}

func (g Token) ChangeId() {
	var id bson.ObjectId = bson.NewObjectId()
	g.ID = &id
}

func (g Token) ChangeUI(ui string) {
	g.UI = ui
}

func (g Token) ChangeGI(gi string) {
	g.GI = gi
}

func (g Token) ChangeRights(rights int32) {
	g.Rights = rights
}

func (g Token) WithDefaultRights() *Token {
	g.Rights = core.DEFAULT_RIGHTS.Value
	return &g
}

func (g Token) GetRights() int32 {
	return g.Rights
}

func (g Token) ChangeCreated() {
	g.Created = time.Now()
}

func (g Token) Value() interface{} {
	return g
}
