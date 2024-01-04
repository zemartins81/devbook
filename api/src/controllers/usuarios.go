package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CriarUsuario cria um novo usuário no sistema.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()
	

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioId, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioId)))
}

// BuscarUsuarios retorna uma lista de todos os usuários.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Implementação da função
}

// BuscarUsuario procura por um usuário específico com base em um identificador.
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
