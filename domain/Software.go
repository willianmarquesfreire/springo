package domain

type Software struct {
	Name          string `json:"name"`
	Url           string `json:"url"`
	GenericDomain        `bson:",inline"`
	Main          bool   `json:"main"`
	TokenDuration int64  `json:"tokenDuration,omitempty"`
}