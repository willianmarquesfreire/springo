package resources

import (
	"springo/domain"
	"springo/rest"
)

type SoftwareServiceType struct {
	rest.Service
}

var SoftwareService = SoftwareServiceType{
	rest.Service{
		Domain:   domain.Software{},
		Document: "softwares",
	},
}
