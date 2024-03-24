package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
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

// PararDeSeguirUsuario chama a api para parar de seguir um usuário
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/parardeseguir", config.ApiUrl, usuarioID)
	response, erro := requisicoes.RequisicoesComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.ApiUrl, usuarioID)
	response, erro := requisicoes.RequisicoesComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// EditarUSuario chama a api para editar um usuário
func EditarUsuario(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

    fmt.Println(string(usuario))

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioID)

	response, err := requisicoes.RequisicoesComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
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

// AtualizarSenha chama a api para atualizar a senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    senhas, err := json.Marshal(map[string]string{
        "atual": r.FormValue("senhaAtual"),
        "nova":  r.FormValue("novaSenha"),
    })
    if err != nil {
        respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
        return
    }

    cookie, _ := cookies.Ler(r)
    usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

    url := fmt.Sprintf("%s/usuario/%d/atualizar-senha", config.ApiUrl, usuarioID)

    fmt.Println(string(senhas))
    fmt.Println(url)

    response, err := requisicoes.RequisicoesComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhas))
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

//DeletarUsuario chama a api para deletar um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
    cookie, _ := cookies.Ler(r)
    usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

    url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioID)

    fmt.Println(url)

    response, err := requisicoes.RequisicoesComAutenticacao(r, http.MethodDelete, url, nil)
    if err != nil {
        respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
        return
    }
    defer response.Body.Close()

    fmt.Println(response.StatusCode)

    if response.StatusCode >= 400 {
        respostas.TratarStatusCodeDeErro(w, response)
        return
    }

    respostas.JSON(w, response.StatusCode, nil)

}
