package resources

import (
	"os"
	"log"
	"io/ioutil"
	"springo/rest"
)

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
	img, err := os.Open("./files/teste.png")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	a, _ := ioutil.ReadAll(img)
	return a

}

func TestandoPdf(w rest.Response, r *rest.Request) []byte {
	file, err := os.Open("./files/teste.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/pdf")
	a, _ := ioutil.ReadAll(file)
	return a
}

func TestandoZip(w rest.Response, r *rest.Request) []byte {
	file, err := os.Open("./files/teste.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/zip")
	a, _ := ioutil.ReadAll(file)
	return a
}