package rotas

import (
	"net/http"
	"webapp/src/controllers"
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
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/buscar-usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioID}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioID}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDePerfilDoUsuarioLogado,
		RequerAutenticacao: true, 
	},
{
        URI:                "/editar-usuario",
        Metodo:             http.MethodGet,
        Funcao:             controllers.CarregarPaginaDeEdicaoDeUsuario,
        RequerAutenticacao: true,
    },
    {
        URI:                "/editar-usuario",
        Metodo:             http.MethodPut,
        Funcao:             controllers.EditarUsuario,
        RequerAutenticacao: true,
    },

    {
        URI:                "/atualizar-senha",
        Metodo:             http.MethodGet,
        Funcao:             controllers.CarregarPaginaDeAtualizacaoDeSenha,
        RequerAutenticacao: true,
    },
    {
        URI:                "/atualizar-senha",
        Metodo:             http.MethodPost,
        Funcao:             controllers.AtualizarSenha,
        RequerAutenticacao: true,
    }, 
{
        URI:                "/deletar-usuario",
        Metodo:             http.MethodDelete, 
        Funcao:             controllers.DeletarUsuario,
        RequerAutenticacao: true,
    },
}
