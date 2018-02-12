package resources

import (
	"springo/domain"
	"springo/rest"
)

var TokenResource = rest.RestResource{
	Path:    "token",
	Service: &TokenService.Service,
	Domain:  domain.Token{},
}