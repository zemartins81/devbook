package rotas

import (
	"net/http"

	"webapp/src/controllers"
)

var rotasPublicacoes = []Rota{

	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
{
        URI:                "/publicacoes/{publicacaoId}/curtir",
        Metodo:             http.MethodGet,
        Funcao:             controllers.CurtirPublicacao,
        RequerAutenticacao: true,
    },
}
