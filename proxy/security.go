package proxy

import (
	"springo/config"
	"net/http"
	"fmt"
)

const Authentication = "/api/public/authenticate"

func Authenticate(token string) (*http.Response, error) {
	fmt.Println(config.MainConfiguration.SecurityUrl + Authentication + "/" + config.MainConfiguration.Identifier + "/" + token)
	return http.Get(config.MainConfiguration.SecurityUrl + Authentication + "/" + config.MainConfiguration.Identifier + "/" + token)
}
