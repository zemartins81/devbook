package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zemartins81/devbookWebApp/src/config"
	"github.com/zemartins81/devbookWebApp/src/cookies"
	"github.com/zemartins81/devbookWebApp/src/router"
	"github.com/zemartins81/devbookWebApp/src/utils"
)

//	func init() {
//		config.HashKey
//	}
func main() {

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()

	r := router.Gerar()
	fmt.Printf("Rodando o WEB APP na Porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
