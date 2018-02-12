package rest

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"springo/core"
	"reflect"
	"springo/domain"
	"log"
	"errors"
	"springo/config"
	"springo/logger"
	"springo/util/mongo"
)

type Result struct {
	PageSize int         `json:"pageSize"`
	Start    int         `json:"start"`
	Count    int         `json:"count"`
	HasNext  int         `json:"hasNext"`
	Values   interface{} `json:"values"`
}

type Search struct {
	PageSize int
	Start    int
	M        bson.M
}

type Service struct {
	Domain   interface{}
	Document string
}

func (service Service) FindAll(search Search) (Result, error) {
	session := Session.Copy()
	defer session.Close()

	c := session.DB(config.MainConfiguration.Database).C(service.Document)

	my := service.Domain
	myType := reflect.TypeOf(my)
	slice := reflect.MakeSlice(reflect.SliceOf(myType), 1, 1)
	values := reflect.New(slice.Type())
	values.Elem().Set(slice)
	count, _ := c.Find(search.M).Count()
	error := c.Find(search.M).Skip(search.Start).Limit(search.PageSize).All(values.Interface())

	response := Result{
		Values:   values.Elem().Interface(),
		PageSize: search.PageSize,
		Start:    search.Start,
		Count:    count,
	}
	return response, error
}

func (service Service) Insert(value domain.GenericInterface) (domain.GenericInterface, error) {
	session := Session.Copy()
	defer session.Close()

	c := session.DB(config.MainConfiguration.Database).C(service.Document)
	value.ChangeId()
	if value.GetRights() == 0 {
		value.ChangeRights(core.DEFAULT_RIGHTS.Value)
	}
	value.ChangeCreated()

	error := c.Insert(&value)
	if error != nil {
		log.Fatalln(error)
	}
	return value, error

}

func (service Service) Update(id string, value domain.GenericInterface) (domain.GenericInterface, error) {
	session := Session.Copy()
	defer session.Close()
	c := session.DB(config.MainConfiguration.Database).C(service.Document)

	error := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &value)
	return value, error
}

func (service Service) Set(id string, value domain.GenericInterface) (domain.GenericInterface, error) {
	session := Session.Copy()
	defer session.Close()
	c := session.DB(config.MainConfiguration.Database).C(service.Document)
	error := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": mongo.GetBsonMSet(value)})
	updated, error := service.Find(id)
	value = updated.(domain.GenericInterface)
	return value, error
}

func (service Service) Collection() (*mgo.Collection) {
	session := Session.Copy()
	defer session.Close()

	return session.DB(config.MainConfiguration.Database).C(service.Document)
}

func (service Service) Find(id string) (interface{}, error) {

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New(logger.ExceptionInvalidId)
	}
	session := Session.Copy()
	defer session.Close()

	c := session.DB(config.MainConfiguration.Database).C(service.Document)

	valueType := reflect.New(reflect.TypeOf(service.Domain))
	value := valueType
	error := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(value.Interface())
	return value.Elem().Interface(), error
}

func (service Service) CreateDatabase() (error) {
	if !config.MainConfiguration.Debug {
		return errors.New("Debug not activated. In production...")
	}

	session := Session.Copy()
	defer session.Close()
	return nil
}

func (service Service) DropDatabase() (error) {
	if !config.MainConfiguration.Debug {
		return errors.New("Debug not activated. In production...")
	}

	session := Session.Copy()
	defer session.Close()
	return session.DB(config.MainConfiguration.Database).DropDatabase()
}

func (service Service) Drop(id string) (string, error) {
	session := Session.Copy()
	defer session.Close()
	c := session.DB(config.MainConfiguration.Database).C(service.Document)

	error := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return id, error
}

func (service Service) EnsureIndex() {
	session := Session.Copy()
	defer session.Close()

	c := session.DB(config.MainConfiguration.Database).C(service.Document)

	index := mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
