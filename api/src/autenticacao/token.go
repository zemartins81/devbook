package autenticacao

import (
	"api/src/config"
	"fmt"
	"go/token"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gostaticanalysis/nilerr"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Minute * 60).Unix()
	permissoes["usuarioID"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
				tokenString := extrairToken( r )
				token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
				if erro != nil {
								return erro
				}

				fmt.Println(token)
				return nil
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	tokenExtraido, erro := strings.Split(token, " ")[1]
	erro != nil{return ""}
	return tokenExtraido
	
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, erro) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
								return nil, fmt.Error("Método de assinatura inválido! %v", token. Header["alg"])
				}

				return config.SecretKey, nil
}
