package main

import (
	"net/http"
	"springo/rest"
	"springo/rest/test/resources"
	"springo/config"
)

func main() {
	Start()
}

func Start() {
	config.StartConfigurationTestMock()
	session := rest.StartSession()
	defer session.Close()
	rest.CurrentApi.OnNotFound(NotFound)
	rest.CurrentApi.Register(rest.Resource{Path: "parametro-nome/{nome}", Method: "GET", Function: resources.ParametroNome})
	rest.CurrentApi.Register(rest.Resource{Path: "parametro-regex/\\d", Method: "GET", Function: resources.ParametroRegex})
	rest.CurrentApi.Register(rest.Resource{Path: "testando-endereco", Method: "GET", Function: resources.TestandoEndereco})
	rest.CurrentApi.Register(rest.Resource{Path: "testando-png", Method: "GET", Function: resources.TestandoPng})
	rest.CurrentApi.Register(rest.Resource{Path: "testando-pdf", Method: "GET", Function: resources.TestandoPdf})
	rest.CurrentApi.Register(rest.Resource{Path: "testando-zip", Method: "GET", Function: resources.TestandoZip})
	rest.CurrentApi.Register(rest.Resource{Path: "public/register", Method: "POST", Function: resources.Register})
	rest.CurrentApi.Register(rest.Resource{Path: "public/authorize/{app}/{token}", Method: "GET", Function: resources.Authorize})
	rest.CurrentApi.RegisterAll(resources.UserResource.GetResources())
	rest.CurrentApi.RegisterAll(resources.TokenResource.GetResources())
	rest.CurrentApi.ListenAndServe()
}

func NotFound(w http.ResponseWriter, r *rest.Request) string {
	return "Rota não encontrada"
}
