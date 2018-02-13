package resources

import (
	"os"
	"log"
	"io/ioutil"
	"springo/rest"
	"springo/domain"
	"strings"
	"net/http"
	"gopkg.in/mgo.v2/bson"
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

func Authenticate(w rest.Response, r *rest.Request) (*domain.Token) {
	if strings.Contains(r.PathVariables["token"].(string), "bloquea") {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}
	return domain.Token{
		Information: bson.NewObjectId().String(),
		Group:domain.Group{
			Software:&domain.Software{
				Name:r.PathVariables["app"].(string),
				Url:r.PathVariables["app"].(string),
			},
			Name:"app",
		}.WithDefaultRights(),
		User:domain.User{
			Login:"willianmarquesfreire@gmail.com",
		}.WithDefaultRights(),
	}.WithDefaultRights()
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