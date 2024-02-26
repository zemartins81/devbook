package rotas

import (
	"net/http"

	"github.com/zemartins81/devbookWebApp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaCriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
}
