package proxy

import (
	"springo/config"
	"net/http"
	"fmt"
)

const Authorization = "/api/public/authorize"
const Authentication = "/api/public/authenticate"

func Authorize(token string) (*http.Response, error) {
	fmt.Println(config.MainConfiguration.SecurityUrl + Authorization + "/" + config.MainConfiguration.Identifier + "/" + token)
	return http.Get(config.MainConfiguration.SecurityUrl + Authorization + "/" + config.MainConfiguration.Identifier + "/" + token)
}
