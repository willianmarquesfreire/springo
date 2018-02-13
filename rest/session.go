package rest

import (
	"springo/config"
	"gopkg.in/mgo.v2"
	"springo/domain"
	"strings"
	"springo/proxy"
	"encoding/json"
	"net/http"
	"errors"
)

var (
	Session *mgo.Session
	err     error
	CurrentApi     *Api
)

func Handler(ww Response, rr *Request, resource Resource) error {
	wtoken := GetParam(rr, "wtoken")

	if strings.Contains(rr.URL.Path, "/public") {
		return nil
	}

	resp,err := proxy.Authenticate(wtoken)
	if resp.StatusCode == http.StatusForbidden {
		ww.Header().Set("Content-Type", "application/json; charset=utf-8")
		ww.WriteHeader(http.StatusForbidden)
		return errors.New("Forbidden")
	}
	var token domain.Token
	err = json.NewDecoder(resp.Body).Decode(&token)

	if resource.Info.RestResource != nil && resource.Info.RestResource.Service != nil{
		resource.Info.RestResource.Service.TokenScope = token
	}

	//ww.WriteHeader(http.StatusBadRequest)
	return err
}

func StartSession() *mgo.Session{
	config.StartConfigurationTestMock()
	CurrentApi = new(Api)
	CurrentApi.NewMux()
	CurrentApi.AddHandler(Handler)
	CurrentApi.Addr = config.MainConfiguration.Addr
	CurrentApi.BaseUrl = config.MainConfiguration.BaseUrl

	Session, err = mgo.Dial(config.MainConfiguration.MgoDial)
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	return Session
}
