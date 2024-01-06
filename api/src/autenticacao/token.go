package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CriarToken gera um token JWT para o ID de usuário fornecido.
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 1).Unix()
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}
