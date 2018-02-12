package resources

import (
	"springo/domain"
	"springo/rest"
)

type UserServiceType struct {
	rest.Service
}

var UserService = UserServiceType{
	rest.Service{
		Domain:   domain.User{},
		Document: "users",
	},
}
