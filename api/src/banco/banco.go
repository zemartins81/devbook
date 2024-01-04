package banco

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Conectar estabelece uma conexão com o banco de dados.
// Retorna um ponteiro para o objeto sql.DB e um erro, se houver.
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
