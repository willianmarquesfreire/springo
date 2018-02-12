package domain

import "time"

type Token struct {
	GenericDomain `json:",omitempty" bson:",inline"`
	Information  string `json:"information"`
	ExpirateAt   time.Time `json:"expirateAt,omitempty"`
	User         *User `json:"user,omitempty"`
	Group     *Group `json:"group,omitempty"`
}