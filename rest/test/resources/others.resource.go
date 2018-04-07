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
	"fmt"
	"encoding/json"
	"reflect"
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

func Authorize(w rest.Response, r *rest.Request) (*domain.Token) {
	if strings.Contains(r.PathVariables["token"].(string), "bloquea") {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}

	id := bson.ObjectIdHex("5a83005233213a5720ef69bd")
	id2 := bson.ObjectIdHex("5a85c98233213a3b116dd8d1")
	u1 := &domain.Token{
		Information: bson.NewObjectId().String(),
		User: &domain.User{
			Login: "willianmarquesfreire@gmail.com",
			GenericDomain: domain.GenericDomain{
				ID: &id,
				GI: "grupowillian,",
			},
		},
	}

	u2 := &domain.Token{
		Information: bson.NewObjectId().String(),
		User: &domain.User{
			Login: "teste@gmail.com",
			GenericDomain: domain.GenericDomain{
				ID: &id2,
				GI: "grupoteste,grupowillian,",
			},
		},
	}

	if r.PathVariables["token"].(string) == "willian" {
		return u1
	}

	return u2
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

func Register(w rest.Response, r *rest.Request) (*domain.Token) {
	valueType := reflect.New(reflect.TypeOf(domain.User{}))
	decoder := json.NewDecoder(r.Body)
	var value domain.GenericInterface = valueType.Interface().(domain.GenericInterface)
	decoder.Decode(value)
	fmt.Println(value, reflect.TypeOf(value))

	uService := UserService
	uResult := uService.SimpleInsert(value).(*domain.User)

	tService := TokenService
	tResult := tService.SimpleInsert(&domain.Token{User:uResult}).(*domain.Token)

	return tResult
}
