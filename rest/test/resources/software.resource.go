package resources

import (
	"springo/domain"
	"springo/rest"
)

var SoftwareResource = rest.RestResource{
	Path:    "software",
	Service: &SoftwareService.Service,
	Domain:  domain.Software{},
}
