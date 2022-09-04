package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
	tokenString := extrairToken(r)
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
	if erro != nil {
		return " "
	}
	return tokenExtraido

}

func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Clains.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint( fmt.Sprintf("%.0f", ["usuarioID"], 10, 64))
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}

	return 0, errors.New("Token unválido")

}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, erro) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Error("Método de assinatura inválido! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
