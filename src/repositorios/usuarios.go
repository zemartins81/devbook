package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
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

// Buscar retorna todos os usuarios que atendem a um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%
	fmt.Println("teste")

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email  from usuarios where nome LIKE ? or nick LIKE ?", nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			//usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		fmt.Println(usuario)

		usuarios = append(usuarios, usuario)
	}

	fmt.Println(usuarios)

	return usuarios, nil

}

func (repositorio Usuarios) BuscarPorId(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(" select id, nome, nick, email, CriadoEM from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario), error {
				statement, erro := repositorio.db.Prepare("update usuarios set nome = ?", nick = )
}
