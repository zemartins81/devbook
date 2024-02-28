package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zemartins81/devbookWebApp/src/utils"
)

// FazerLogin faz a autenticacao do usuario
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string){
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	response, err := http.Post(
		"http://localhost:5000/login",
		"application/json",
		bytes.NewBuffer(usuario),
	)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	fmt.Println(response.StatusCode, response.Body)
}
