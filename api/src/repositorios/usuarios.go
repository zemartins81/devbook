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

// Atualizar atualiza as informações de um usuário no banco de dados.
func (u Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := u.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		ID,
	); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de Dados
func (u Usuarios) Deletar(ID uint64) error {
	statement, erro := u.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (u Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := u.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (u Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := u.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil

}

func (u Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := u.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (u Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := u.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm
	from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?`,
		usuarioID)
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
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}



	return usuarios, nil
}

func (u Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := u.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.CriadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios = []modelos.Usuario{}

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	
    fmt.Println(usuarios)
	return usuarios, nil 

}

func (u Usuarios) BuscarSenha(id uint64) (string, error) {
	linha, erro := u.db.Query("select senha from usuarios where id = ?", id)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Senha

	if linha.Next() {
		if erro = linha.Scan(&usuario.Atual); erro != nil {
			return "", erro
		}
	}

	return usuario.Atual, nil
}

func (u Usuarios) AtualizarSenha(id uint64, senha string) error {
	statement, erro := u.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, id); erro != nil {
		return erro
	}

	return nil
}
