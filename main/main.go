package main

import (
	"net/http"
	"fmt"
	"reflect"
	"encoding/json"
	"strings"
	"regexp"
)

type Pessoa struct {
	Nome string `json:"nome"`
}

type Api struct {
	resources map[string][]Resource
	Addr      string
	NotFound  interface{}
	BaseUrl   string
}

func (api Api) NewMux() Api {
	api.resources = make(map[string][]Resource)
	return api
}

func (api Api) ListenAndServe() Api {

	http.HandleFunc(api.BaseUrl, func(ww http.ResponseWriter, rr *http.Request) {
		fmt.Println(rr.URL.Path)
		args := make([]reflect.Value, 2)
		args[0] = reflect.ValueOf(ww)
		request := &Request{}
		request.Request = rr
		args[1] = reflect.ValueOf(request)
		var values []reflect.Value

		rs := api.resources[strings.Split(rr.URL.Path, "/")[0]]
		var r Resource

		for _, element := range rs {
			if len(element.Info.Regex.FindString(rr.URL.Path)) > 0 {
				r = element
				break
			}
		}

		request.PathVariables = make(map[string]interface{})
		splitUrl := strings.Split(rr.URL.Path, "/")
		for key, element := range r.Info.PathVariables {
			request.PathVariables[key] = splitUrl[element]
		}

		if r.Method == rr.Method {
			values = reflect.ValueOf(r.Function).Call(args)
		} else {
			values = reflect.ValueOf(api.NotFound).Call(args)
		}

		var respBody []byte

		if values != nil && len(values) > 0 {
			respBody, _ = json.MarshalIndent(values[0].Interface(), "", "  ")
		} else {
			respBody, _ = json.MarshalIndent(values, "", "  ")
		}
		ww.Header().Set("Content-Type", "application/json; charset=utf-8")
		ww.Write(respBody)
	})

	http.ListenAndServe(api.Addr, nil)
	return api
}

func (api *Api) OnNotFound(fnNotFound interface{}) *Api {
	api.NotFound = fnNotFound
	return api
}

func (api *Api) Register(r Resource) *Api {
	bars := strings.Split(r.Path, "/")
	r.Info = ResourceInfo{}
	r.Info.PathVariables = make(map[string]int)

	for i, elem := range bars {
		if len(regexp.MustCompile("{\\w+}").FindString(elem)) > 0 {
			r.Info.PathVariables[regexp.MustCompile("{|}").ReplaceAllString(elem, "")] = i
		}
	}

	r.Info.Regex, _ = regexp.Compile("^" + regexp.MustCompile("{\\w+}").ReplaceAllString(r.Path, "(\\w+)") + "$")
	pos := strings.Split(r.Path, "/")[0]
	api.resources[pos] = append(api.resources[pos], r)

	return api
}

type ResourceInfo struct {
	PathVariables map[string]int
	Regex         *regexp.Regexp
}

type Resource struct {
	Path     string
	Method   string
	Function interface{}
	Info     ResourceInfo
}

type Request struct {
	*http.Request
	PathVariables map[string]interface{}
}

func Testando(w http.ResponseWriter, r *Request) string {
	fmt.Println(r.PathVariables)
	return "{\"zxczxc\":\"\asd\"}"
	//return Pessoa{Nome: "will"}
}

func NotFound(w http.ResponseWriter, r *Request) string {
	return "erro"
}

func main() {
	api := Api{Addr: ":8080", BaseUrl: "/"}.NewMux()
	api.OnNotFound(NotFound)
	api.Register(Resource{Path: "/asd/{bbb}", Method: "GET", Function: Testando})
	api.ListenAndServe()
	//http.HandleFunc("/teste", Testando)
}
