package main

import (
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"springo/rest"
	"springo/config"
)

func main() {
	config.StartConfigurationTestMock()
	api := rest.Api{Addr: ":8086", BaseUrl: "/api/"}.NewMux()
	api.OnNotFound(NotFound)
	api.Register(rest.Resource{Path: "testando/{nome}", Method: "GET", Function: Testando})
	api.Register(rest.Resource{Path: "testando1/\\d", Method: "GET", Function: Testando1})
	api.Register(rest.Resource{Path: "testando2", Method: "GET", Function: Testando2})
	api.Register(rest.Resource{Path: "testando3", Method: "GET", Function: Testando3})
	api.Register(rest.Resource{Path: "testando4", Method: "GET", Function: Testando4})
	api.Register(rest.Resource{Path: "testando5", Method: "GET", Function: Testando5})
	api.ListenAndServe()
}

type Pessoa struct {
	Nome string `json:"nome"`
}

type Endereco struct {
	Rua string `json:"rua"`
}

func Testando(w rest.Response, r *rest.Request) Pessoa {
	a := r.PathVariables["nome"]
	return Pessoa{Nome: a.(string)}
}

func Testando1(w rest.Response, r *rest.Request) string {
	return "Willian"
}


func Testando2(w rest.Response, r *rest.Request) Endereco {
	return Endereco{Rua: "Rua Willian"}
}

func Testando3(w rest.Response, r *rest.Request) []byte {
	img, err := os.Open("./teste.png")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	a, _ := ioutil.ReadAll(img)
	return a

}

func Testando4(w rest.Response, r *rest.Request) []byte {
	file, err := os.Open("./teste.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/pdf")
	a, _ := ioutil.ReadAll(file)
	return a
}

func Testando5(w rest.Response, r *rest.Request) []byte {
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

