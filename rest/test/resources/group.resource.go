package resources

import (
	"springo/domain"
	"springo/rest"
)

var GroupResource = rest.RestResource{
	Path:    "group",
	Service: &GroupService.Service,
	Domain:  domain.Group{},
}
