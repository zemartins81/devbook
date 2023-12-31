package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar cria e retorna uma nova instância de mux.Router.
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
