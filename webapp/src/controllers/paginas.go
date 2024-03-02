package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zemartins81/devbookWebApp/src/config"
	"github.com/zemartins81/devbookWebApp/src/modelos"
	"github.com/zemartins81/devbookWebApp/src/requisicoes"
	"github.com/zemartins81/devbookWebApp/src/respostas"
	"github.com/zemartins81/devbookWebApp/src/utils"
)

// CarregarTelaDeLogin Carrega a página de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaCriarUsuario Carrega a página de cadastro
func CarregarPaginaCriarUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal Carrega a página principal
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.ApiUrl)
	response, err := requisicoes.RequisicoesComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "home.html", publicacoes)
}
