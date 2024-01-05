package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
	"time"
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

// Buscar encontra usuários pelo nome ou apelido.
// Retorna um slice de structs Usuario e um erro, se houver.
func (u Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		var criadoEmBytes []byte
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&criadoEmBytes,
		); erro != nil {
			return nil, erro
		}
		criadoEmStr := string(criadoEmBytes)
		usuario.CriadoEm, erro = time.Parse("2006-01-02 15:04:05", criadoEmStr) // Use o formato correto
		if erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

// BuscarPorID busca um usuário pelo seu ID.
// Retorna o usuário encontrado ou um erro, caso não seja encontrado.
func (u Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := u.db.Query(
		"select id, nome, nick, email, criadoEM from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	var criadoEmBytes []byte
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&criadoEmBytes,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
		criadoEmStr := string(criadoEmBytes)
		usuario.CriadoEm, erro = time.Parse("2006-01-02 15:04:05", criadoEmStr) // Use o formato correto
		if erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}
