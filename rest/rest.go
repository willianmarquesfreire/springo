package rest

import (
	"regexp"
	"net/http"
	"reflect"
	"encoding/json"
	"strings"
	"springo/logger"
	"springo/config"
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
func (api *Api) NewMux() *Api {
	api.resources = make(map[string][]Resource)
	return api
}

/**
	Inicializa a API de acordo com as configurações realizadas
 */
func (api Api) ListenAndServe() Api {

	baseUrl := api.BaseUrl
	if baseUrl == "" {
		baseUrl = "/"
	}

	logger.MessageApiStartedLog(config.MainConfiguration.ApiPath+baseUrl, api.Addr)
	http.HandleFunc(config.MainConfiguration.ApiPath+baseUrl, func(ww http.ResponseWriter, rr *http.Request) {
		respBody := mountResponseJSON(ww, rr.WithContext(rr.Context()), api)
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

	err := http.ListenAndServe(api.Addr, nil)
	if err != nil {
		logger.ExceptionApiNotStartedLog(err.Error())
	}
	return api
}

/**
	Monta JSON de resposta
 */
func mountResponseJSON(ww http.ResponseWriter, rr *http.Request, api Api) interface{} {
	request := &Request{}
	request.Request = rr

	response := Response{}
	response.ResponseWriter = ww

	args := mountArgsListAccordingParamsToCallMethod(response, request)
	resource := getResourceOnApiOfRequest(api, request)
	request.PathVariables = getPathVariablesOfRequestOnResource(request, resource)
	values := getReturningValuesOfCallMethodOfResourceAccordingParams(resource, request, args, api)

	respBody := mountRespBodyAccordingValues(values)
	return respBody
}

/**
	Função que monta um json de acordo com a lista de valores. Utilizada na função MountResponseJSON
 */
func mountRespBodyAccordingValues(values []reflect.Value) interface{} {
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
func getReturningValuesOfCallMethodOfResourceAccordingParams(resource Resource, rr *Request, args []reflect.Value, api Api) []reflect.Value {
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
func mountArgsListAccordingParamsToCallMethod(response Response, request *Request) []reflect.Value {
	args := make([]reflect.Value, 2)
	args[0] = reflect.ValueOf(response)
	args[1] = reflect.ValueOf(request)
	return args
}

/**
	Retorna Lista de variáveis de URL de acordo com a requisição e recurso
 */
func getPathVariablesOfRequestOnResource(rr *Request, resource Resource) map[string]interface{} {
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
func getResourceOnApiOfRequest(api Api, rr *Request) Resource {
	baseUrl := api.BaseUrl
	if baseUrl == "" {
		baseUrl = "/"
	}
	reg := regexp.MustCompile(config.MainConfiguration.ApiPath + baseUrl).ReplaceAllString(rr.URL.Path, "")

	rs := api.resources[strings.Split(reg, "/")[0]]
	var r Resource
	for _, element := range rs {
		if len(element.Info.Regex.FindString(reg)) > 0 {
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


	baseUrl := api.BaseUrl
	if baseUrl == "" {
		baseUrl = "/"
	}

	bars := strings.Split(baseUrl+r.Path, "/")
	r.Info = ResourceInfo{}
	r.Info.PathVariables = make(map[string]int)

	for i, elem := range bars {
		if len(regexp.MustCompile("{\\w+}").FindString(elem)) > 0 {
			r.Info.PathVariables[regexp.MustCompile("{|}").ReplaceAllString(elem, "")] = i + 1
		}
	}

	r.Info.Regex, _ = regexp.Compile("^" + regexp.MustCompile("{\\w+}").ReplaceAllString(r.Path, "(\\w+)") + "$")
	pos := strings.Split(r.Path, "/")[0]

	api.resources[pos] = append(api.resources[pos], r)

	logger.MessageResourceStartedLog(r.Path)

	return api
}

func (api *Api) RegisterAll(r []Resource) *Api {
	for _, elm := range r {
		api.Register(elm)
	}
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
