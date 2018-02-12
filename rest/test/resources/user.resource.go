package resources

import (
	"springo/domain"
	"springo/rest"
)

var UserResource = rest.RestResource{
	Path:    "user",
	Service: &UserService.Service,
	Domain:  domain.User{},
}
