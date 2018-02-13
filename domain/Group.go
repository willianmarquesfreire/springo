package domain

import (
	"time"
	"springo/core"
	"gopkg.in/mgo.v2/bson"
)

type Group struct {
	GenericDomain `bson:",inline"`
	Name           string `json:"name"`
	SocialName     string `json:"socialName"`
	SoftwareHouse  bool `json:"softwareHouse"`
	Expiration        time.Time `json:"expiration"`
	MaxShortTokens int `json:"maxShortTokens"`
	CustomInformation string `json:"customInformation"`
	TokenDuration     int64 `json:"tokenDuration"`
	Software *Software `json:"software,omitempty"`
}

func (g Group) ChangeId() {
	var id bson.ObjectId = bson.NewObjectId()
	g.ID = &id
}

func (g Group) ChangeUI(ui string) {
	g.UI = ui
}

func (g Group) ChangeGI(gi string) {
	g.GI = gi
}

func (g Group) ChangeRights(rights int32) {
	g.Rights = rights
}

func (g Group) WithDefaultRights() *Group {
	g.Rights = core.DEFAULT_RIGHTS.Value
	return &g
}

func (g Group) GetRights() int32 {
	return g.Rights
}

func (g Group) ChangeCreated() {
	g.Created = time.Now()
}

func (g Group) Value() interface{} {
	return g
}
