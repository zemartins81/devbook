package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarUsuario cria um novo usuário no sistema.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios retorna uma lista de todos os usuários.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// Implementação da função
}

// AtualizarUsuario atualiza os dados de um usuário existente.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// Implementação da função
}

// DeletarUsuario remove um usuário do sistema.
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	// Implementação da função
}
