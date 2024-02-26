package controllers

import (
	"net/http"

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
