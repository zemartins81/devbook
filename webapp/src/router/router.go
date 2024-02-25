package router

import (
	"github.com/gorilla/mux"
	"github.com/zemartins81/devbookWebApp/src/router/rotas"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	rotas.Configurar(r)
	return r
}
