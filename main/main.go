package main

import (
	"net/http"
	"fmt"
	"reflect"
	"encoding/json"
	"strings"
	"regexp"
	"os"
	"log"
	"io"
	"io/ioutil"
)

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
		respBody := MountResponseJSON(ww, rr, api)
		//ww.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Println(respBody, reflect.TypeOf(respBody))
		if reflect.TypeOf(respBody).String() != "[]uint8" {
			respBody, _ = json.MarshalIndent(respBody, "", "  ")
			ww.Write(respBody.([]byte))
		} else {
			ww.Write(respBody.([]uint8))
		}
	})

	http.ListenAndServe(api.Addr, nil)
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

type Response struct {
	http.ResponseWriter
}

type Pessoa struct {
	Nome string `json:"nome"`
}

type Endereco struct {
	RuaLegal string `json:"ruaLegal"`
}

func Testando(w Response, r *Request) Pessoa {
	//return "{\"zxczxc\":\"\asd\"}"
	a := r.PathVariables["bbb"]
	return Pessoa{Nome: a.(string)}
}

func Testando1(w Response, r *Request) string {
	//return "{\"zxczxc\":\"\asd\"}"
	return "Douglas"
}


func Testando2(w Response, r *Request) Endereco {
	//return "{\"zxczxc\":\"\asd\"}"
	return Endereco{RuaLegal : "will"}
}

func Testando3(w Response, r *Request) []byte {
	img, err := os.Open("./teste.png")
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	a, _ := ioutil.ReadAll(img)
	return a

}
func Testando4(w Response, r *Request) string {
	img, err := os.Open("./teste.png")
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)
	return ""
}

func NotFound(w http.ResponseWriter, r *Request) string {
	return "Rota não encontrada"
}

func main() {
	api := Api{Addr: ":8086", BaseUrl: "/"}.NewMux()
	api.OnNotFound(NotFound)
	api.Register(Resource{Path: "/asd/{bbb}", Method: "GET", Function: Testando2})
	api.ListenAndServe()
	//http.HandleFunc("/teste", Testando)
}




//func Home(w http.ResponseWriter, r *http.Request) {
//	imgFile, err := os.Open("./teste.png") // a QR code image
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	defer imgFile.Close()
//
//	// create a new buffer base on file size
//	fInfo, _ := imgFile.Stat()
//	var size int64 = fInfo.Size()
//	buf := make([]byte, size)
//
//	// read file content into buffer
//	fReader := bufio.NewReader(imgFile)
//	fReader.Read(buf)
//
//	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()
//
//	// png.Encode(&buf, image)
//
//	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
//	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
//
//	// Embed into an html without PNG file
//	img2html := "<html><body><img src=\"data:image/png;base64," + imgBase64Str + "\" /></body></html>"
//
//	w.Write([]byte(fmt.Sprintf(img2html)))
//
//}
//
//func main() {
//	// http.Handler
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", Home)
//
//	http.ListenAndServe(":8086", mux)
//}