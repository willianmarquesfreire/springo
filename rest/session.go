package rest

import (
	"springo/config"
	"gopkg.in/mgo.v2"
	"springo/domain"
	"net/http"
)

var (
	Session *mgo.Session
	err     error
	CurrentApi     *Api
)

func Handler(ww Response, rr *Request, resource Resource) error {
	let := GetParam(rr, "wtoken")
	if resource.Info.RestResource != nil && resource.Info.RestResource.Service != nil{
		resource.Info.RestResource.Service.TokenScope = domain.Token{
			Information: let,
		}
	}
	ww.Header().Set("Content-Type", "application/json; charset=utf-8")
	ww.WriteHeader(http.StatusBadRequest)
	return nil
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
