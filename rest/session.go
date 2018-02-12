package rest

import (
	"springo/config"
	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	err     error
	CurrentApi     *Api
)

func StartSession() *mgo.Session{
	config.StartConfigurationTestMock()
	CurrentApi = new(Api)
	CurrentApi.Addr = config.MainConfiguration.Addr
	CurrentApi.BaseUrl = config.MainConfiguration.BaseUrl
	CurrentApi.NewMux()

	Session, err = mgo.Dial(config.MainConfiguration.MgoDial)
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	return Session
}
