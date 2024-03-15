package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"webapp/src/config"
	"webapp/src/respostas"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
