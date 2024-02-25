package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zemartins81/devbookWebApp/src/router"
	"github.com/zemartins81/devbookWebApp/src/utils"
)

	func main() {

		utils.CarregarTemplates()

		r := router.Gerar()
		fmt.Println("Rodando o WEB APP")
		log.Fatal(http.ListenAndServe(":3000", r))
	}
