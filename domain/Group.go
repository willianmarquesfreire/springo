package domain

import "time"

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