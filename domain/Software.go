package domain

import (
	"springo/core"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Software struct {
	Name          string `json:"name"`
	Url           string `json:"url"`
	GenericDomain        `bson:",inline"`
	TokenDuration int64  `json:"tokenDuration,omitempty"`
}

func (g Software) ChangeId() {
	var id bson.ObjectId = bson.NewObjectId()
	g.ID = &id
}

func (g Software) ChangeUI(ui string) {
	g.UI = ui
}

func (g Software) ChangeGI(gi string) {
	g.GI = gi
}

func (g Software) ChangeRights(rights int32) {
	g.Rights = rights
}

func (g Software) WithDefaultRights() *Software {
	g.Rights = core.DEFAULT_RIGHTS.Value
	return &g
}

func (g Software) GetRights() int32 {
	return g.Rights
}

func (g Software) ChangeCreated() {
	g.Created = time.Now()
}

func (g Software) Value() interface{} {
	return g
}
