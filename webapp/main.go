package main

import (
	"fmt"
	"log"
	"net/http"

	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"

)

// func init() {
// 	HashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	BlockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(HashKey, BlockKey)
// }

func main() {

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()

	r := router.Gerar()
	fmt.Printf("Rodando o WEB APP na Porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
