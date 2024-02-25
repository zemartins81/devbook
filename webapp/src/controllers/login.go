package controllers

import (
	"net/http"

	"github.com/zemartins81/devbookWebApp/src/utils"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
