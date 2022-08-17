package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) CriarUsuario(usuario modelos.Usuario) (uint64, error) {
	stmt, err := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuario.Senha,
	)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil

}
