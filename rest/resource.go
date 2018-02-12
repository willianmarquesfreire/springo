package rest


import (
	"net/http"
	"gopkg.in/mgo.v2"
	"goji.io"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"reflect"
	"springo/domain"
)

type ResourceInterface interface {
	Init(mux *goji.Mux, session *mgo.Session)
	InitResource(mux *goji.Mux, session *mgo.Session, service *Service)
	ErrorWithJSON(w http.ResponseWriter, message string, code int)
	ResponseWithJSON(w http.ResponseWriter, json []byte, code int)
	GetToken(r *http.Request) string
	NewSearch(r *http.Request) Search
	GetAll() func(w http.ResponseWriter, r *http.Request)
	Post() func(w http.ResponseWriter, r *http.Request)
	Get() func(w http.ResponseWriter, r *http.Request)
	Put() func(w http.ResponseWriter, r *http.Request)
	Delete() func(w http.ResponseWriter, r *http.Request)
}

type RestResource struct {
	Domain    interface{}
	Path      string
	Document  string
	Service   *Service
	Resources []Resource
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *RestResource) GetResources() []Resource{
	resources := []Resource{
		Resource{Path: r.Path, Method: "GET", Function: r.GetAll, Info:ResourceInfo{RestResource:r}},
		Resource{Path: r.Path + "/{id}", Method: "DELETE", Function: r.Delete, Info:ResourceInfo{RestResource:r}},
		Resource{Path: r.Path + "/{id}", Method: "GET", Function: r.Get, Info:ResourceInfo{RestResource:r}},
		Resource{Path: r.Path, Method: "POST", Function: r.Post, Info:ResourceInfo{RestResource:r}},
		Resource{Path: r.Path + "/{id}", Method: "PUT", Function: r.Put, Info:ResourceInfo{RestResource:r}},
		Resource{Path: r.Path + "/{id}", Method: "PATCH", Function: r.Patch, Info:ResourceInfo{RestResource:r}},
	}
	if r.Resources == nil {
		r.Resources = resources
	} else {
		r.Resources = append(r.Resources, resources...)
	}
	return r.Resources
}

func (r *RestResource) AddResource(resource Resource) RestResource {
	resource.Info.RestResource = r
	r.Resources = append(r.Resources, resource)
	return *r
}

func (r RestResource) RegisterOn(api *Api) RestResource {
	api.RegisterAll(r.GetResources())
	return r
}

func GetParam(r *Request, param string) string {
	var resp string
	resp = r.URL.Query().Get(param)
	if resp == "" {
		resp = r.Header.Get(param)
	}
	return resp
}

func NewSearch(r *Request) Search {
	qo := Search{}
	strPageSize := GetParam(r, "pageSize")
	if strPageSize == "" {
		strPageSize = "10"
	}
	strStart := GetParam(r, "start")
	qo.PageSize, _ = strconv.Atoi(strPageSize)
	qo.Start, _ = strconv.Atoi(strStart)

	strM := GetParam(r, "m")
	var M bson.M
	if strM == "" {
		M = bson.M{}
	} else {
		json.Unmarshal([]byte(strM), &M)
	}
	qo.M = M
	return qo
}

func (resource RestResource) GetAll(w Response, r *Request) Result {
	toReturn, _ := resource.Service.FindAll(NewSearch(r))
	return toReturn
}

func (resource RestResource) Post(w Response, r *Request) interface{} {
	value := resource.DecodeBody(r)
	saved, _ := resource.Service.Insert(value)
	return saved
}


func (resource RestResource) Get(w Response, r *Request) interface{} {
	toReturn, _ := resource.Service.Find(r.PathVariables["id"].(string))
	return toReturn
}

func (resource RestResource) Put(w Response, r *Request) interface{} {
	value := resource.DecodeBody(r)
	saved, _ := resource.Service.Update(r.PathVariables["id"].(string), value)
	return saved
}

func (resource RestResource) Patch(w Response, r *Request) interface{} {
	value := resource.DecodeBody(r)
	saved, _ := resource.Service.Set(r.PathVariables["id"].(string), value)
	return saved
}

func (resource RestResource) Delete(w Response, r *Request) string {
	toReturn, _ := resource.Service.Drop(r.PathVariables["id"].(string))
	return toReturn
}

func (resource RestResource) DecodeBody(r *Request) domain.GenericInterface {
	valueType := reflect.New(reflect.TypeOf(resource.Domain))
	decoder := json.NewDecoder(r.Body)
	var value domain.GenericInterface = valueType.Interface().(domain.GenericInterface)
	decoder.Decode(value)
	return value
}