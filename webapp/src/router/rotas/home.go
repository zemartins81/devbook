package rotas

import (
	"net/http"

	"github.com/zemartins81/devbookWebApp/src/controllers"
)

var rotaHome = Rota{

	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
