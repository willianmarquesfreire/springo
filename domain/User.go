package domain

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"springo/core"
)

type User struct {
	GenericDomain      `bson:",inline"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

func (g User) ChangeId() {
	var id bson.ObjectId = bson.NewObjectId()
	g.ID = &id
}

func (g User) ChangeUI(ui string) {
	g.UI = ui
}

func (g User) ChangeGI(gi string) {
	g.GI = gi
}

func (g User) ChangeRights(rights int32) {
	g.Rights = rights
}

func (g User) WithDefaultRights() *User {
	g.Rights = core.DEFAULT_RIGHTS.Value
	return &g
}

func (g User) GetRights() int32 {
	return g.Rights
}

func (g User) ChangeCreated() {
	g.Created = time.Now()
}

func (g User) Value() interface{} {
	return g
}
