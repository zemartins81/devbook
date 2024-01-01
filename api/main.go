package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	fmt.Println(config.StringConexaoBanco)

	r := router.Gerar()

	fmt.Println("Escutando na porta 5000")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
