package banco

import (
	"api/src/config"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoDB)
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
