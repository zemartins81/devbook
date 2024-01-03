package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria uma nova instância do repositório de usuários.
// Recebe um *sql.DB como parâmetro e retorna um ponteiro para uma struct usuarios.
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um novo usuário no banco de dados.
// Retorna o ID do usuário criado ou um erro.
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := u.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuario.Senha,
	)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}
