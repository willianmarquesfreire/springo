package resources

import (
	"springo/domain"
	"springo/rest"
)

type GroupServiceType struct {
	rest.Service
}

var GroupService = GroupServiceType{
	rest.Service{
		Domain:   domain.Group{},
		Document: "groups",
	},
}
