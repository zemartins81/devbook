package router

import "github.com/gorilla/mux"


// Gerar cria e retorna uma nova instância de mux.Router.
func Gerar() *mux.Router {
	return mux.NewRouter()
}

