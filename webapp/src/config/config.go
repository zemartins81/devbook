package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/zemartins81/devbookWebApp/src/secret"
)

var (
	// ApiUrl e a URL do servidor API
	ApiUrl = "http://localhost:5000"
	// Porta onde a aplicacão web vai rodar
	Porta = 0
	// Chave de autenticação do Cookie
	HashKey = []byte("")
	// Chave de segurança do Cookie
	BlockKey = []byte("")
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	ApiUrl = os.Getenv("API_URL")
	HashKey = secret.GerarSecretKey()
	BlockKey = secret.GerarSecretKey()

}
