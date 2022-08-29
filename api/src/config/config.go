package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	// StringConexaoDB é a string de conexão com o DB
	StringConexaoDB = ""
	Porta           = 0
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 5000
	}

	StringConexaoDB = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

}
