package domain

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	GenericDomain `json:",omitempty" bson:",inline"`
	Information  string `json:"information"`
	ExpirateAt   time.Time `json:"expirateAt,omitempty"`
	User         *User `json:"user,omitempty"`
}

func (g *Token) ChangeId() {
	var id bson.ObjectId = bson.NewObjectId()
	g.ID = &id
}
