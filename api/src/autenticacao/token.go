package autenticacao

import (
	"go/token"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"
)

//CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Minute * 60).Unix()
	permissoes["usuarioID"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("secret"))
}
