package rest

import (
	"regexp"
	"net/http"
	"reflect"
	"encoding/json"
	"strings"
	"springo/logger"
)

/**
	API - Application Programming Interface
 */
type Api struct {
	resources map[string][]Resource
	Addr      string
	NotFound  interface{}
	BaseUrl   string
}

/**
	Inicializa configurações iniciais da API
 */
func (api Api) NewMux() Api {
	api.resources = make(map[string][]Resource)
	return api
}

/**
	Inicializa a API de acordo com as configurações realizadas
 */
func (api Api) ListenAndServe() Api {

	logger.MessageApiStartedLog(api.BaseUrl, api.Addr)
	http.HandleFunc(api.BaseUrl, func(ww http.ResponseWriter, rr *http.Request) {
		respBody := MountResponseJSON(ww, rr, api)
		typeof := reflect.TypeOf(respBody).String()

		if ww.Header().Get("Content-Type") == "" {
			if typeof != "string" {
				ww.Header().Set("Content-Type", "application/json; charset=utf-8")
			}
		}

		if typeof != "[]uint8" {
			respBody, _ = json.MarshalIndent(respBody, "", "  ")
			ww.Write(respBody.([]byte))
		} else {
			ww.Write(respBody.([]uint8))
		}
	})

	error := http.ListenAndServe(api.Addr, nil)
	if error != nil {
		logger.ExceptionApiNotStartedLog(error.Error())
	}
	return api
}

/**
	Monta JSON de resposta
 */
func MountResponseJSON(ww http.ResponseWriter, rr *http.Request, api Api) interface{} {
	request := &Request{}
	request.Request = rr

	response := Response{}
	response.ResponseWriter = ww

	args := MonteArgsListAccordingParamsToCallMethod(response, request)
	resource := GetResourceOnApiOfRequest(api, request)
	request.PathVariables = GetPathVariablesOfRequestOnResource(request, resource)
	values := GetReturningValuesOfCallMethodOfResourceAccordingParams(resource, request, args, api)

	respBody := MountRespBodyAccordingValues(values)
	return respBody
}

/**
	Função que monta um json de acordo com a lista de valores. Utilizada na função MountResponseJSON
 */
func MountRespBodyAccordingValues(values []reflect.Value) interface{} {
	var respBody interface{}
	if values != nil && len(values) > 0 {
		respBody = values[0].Interface()
	} else {
		respBody, _ = json.MarshalIndent(values, "", "  ")
	}
	return respBody
}

/**
	Verifica método da requisição com do recurso utilizado, e chama a função do recurso com determinados argumentos. Se o método do recurso for diferente da requisição, é retornado o método NotFound da api
 */
func GetReturningValuesOfCallMethodOfResourceAccordingParams(resource Resource, rr *Request, args []reflect.Value, api Api) []reflect.Value {
	var values []reflect.Value
	if resource.Method == rr.Method {
		values = reflect.ValueOf(resource.Function).Call(args)
	} else {
		values = reflect.ValueOf(api.NotFound).Call(args)
	}
	return values
}

/**
	Monta lista de argumentos que serão utilizados no método GetReturningValuesOfCallMethodOfResourceAccordingParams
 */
func MonteArgsListAccordingParamsToCallMethod(response Response, request *Request) []reflect.Value {
	args := make([]reflect.Value, 2)
	args[0] = reflect.ValueOf(response)
	args[1] = reflect.ValueOf(request)
	return args
}

/**
	Retorna Lista de variáveis de URL de acordo com a requisição e recurso
 */
func GetPathVariablesOfRequestOnResource(rr *Request, resource Resource) map[string]interface{} {
	splitUrl := strings.Split(rr.URL.Path, "/")
	aux := make(map[string]interface{})
	for key, element := range resource.Info.PathVariables {
		aux[key] = splitUrl[element]
	}
	return aux
}

/**
	Retorna recurso da Api utilizado na requisição
 */
func GetResourceOnApiOfRequest(api Api, rr *Request) Resource {
	rs := api.resources[strings.Split(rr.URL.Path, "/")[0]]
	var r Resource
	for _, element := range rs {
		if len(element.Info.Regex.FindString(rr.URL.Path)) > 0 {
			r = element
			break
		}
	}
	return r
}

/**
	Função executada quando ouver algum erro no mapeamento de determinada rota
 */
func (api *Api) OnNotFound(fnNotFound interface{}) *Api {
	api.NotFound = fnNotFound
	return api
}

/**
	Registra Recurso da API
 */
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

	logger.MessageResourceStartedLog(r.Path)

	return api
}

/**
	Informações adicionais do recurso
 */
type ResourceInfo struct {
	PathVariables map[string]int
	Regex         *regexp.Regexp
}

/**
	Recurso da API
 */
type Resource struct {
	Path     string
	Method   string
	Function interface{}
	Info     ResourceInfo
}

/**
	Requisição
 */
type Request struct {
	*http.Request
	PathVariables map[string]interface{}
}

/**
	Resposta da requisição
 */
type Response struct {
	http.ResponseWriter
}
