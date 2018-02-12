package main

import (
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"springo/rest"
	"springo/domain"
)

type UserServiceType struct {
	rest.Service
}

var UserService = UserServiceType{
	rest.Service{
		Domain:   domain.User{},
		Document: "users",
	},
}

var UserResource = rest.RestResource{
	Path:    "user",
	Service: &UserService.Service,
	Domain: domain.User{},
}

func main() {
	session := rest.StartSession()
	defer session.Close()

	rest.CurrentApi.OnNotFound(NotFound)
	UserResource.RegisterOn(rest.CurrentApi)
	rest.CurrentApi.Register(rest.Resource{Path: "parametro-nome/{nome}", Method: "GET", Function: ParametroNome})
	rest.CurrentApi.Register(rest.Resource{Path: "parametro-regex/\\d", Method: "GET", Function: ParametroRegex})
	rest.CurrentApi.Register(rest.Resource{Path: "testando-endereco", Method: "GET", Function: TestandoEndereco})
	rest.CurrentApi.Register(rest.Resource{Path: "testando-png", Method: "GET", Function: TestandoPng})
	rest.CurrentApi.Register(rest.Resource{Path: "testando-pdf", Method: "GET", Function: TestandoPdf})
	rest.CurrentApi.Register(rest.Resource{Path: "testando-zip", Method: "GET", Function: TestandoZip})
	rest.CurrentApi.ListenAndServe()
}

type Pessoa struct {
	Nome string `json:"nome"`
}

type Endereco struct {
	Rua string `json:"rua"`
}

func ParametroNome(w rest.Response, r *rest.Request) Pessoa {
	a := r.PathVariables["nome"]
	return Pessoa{Nome: a.(string)}
}

func ParametroRegex(w rest.Response, r *rest.Request) string {
	return "Willian"
}

func TestandoEndereco(w rest.Response, r *rest.Request) Endereco {

	return Endereco{Rua: "Rua Willian"}
}

func TestandoPng(w rest.Response, r *rest.Request) []byte {
	img, err := os.Open("./teste.png")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	a, _ := ioutil.ReadAll(img)
	return a

}

func TestandoPdf(w rest.Response, r *rest.Request) []byte {
	file, err := os.Open("./teste.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/pdf")
	a, _ := ioutil.ReadAll(file)
	return a
}

func TestandoZip(w rest.Response, r *rest.Request) []byte {
	file, err := os.Open("./teste.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/zip")
	a, _ := ioutil.ReadAll(file)
	return a
}

func NotFound(w http.ResponseWriter, r *rest.Request) string {
	return "Rota não encontrada"
}
