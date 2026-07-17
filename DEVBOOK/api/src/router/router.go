package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

// A função Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routers.Configurar(r)
}
