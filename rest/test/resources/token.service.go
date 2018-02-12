package resources

import (
	"springo/domain"
	"springo/rest"
)

type TokenServiceType struct {
	rest.Service
}

var TokenService = TokenServiceType{
	rest.Service{
		Domain:   domain.Token{},
		Document: "tokens",
	},
}
