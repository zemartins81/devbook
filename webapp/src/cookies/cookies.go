package cookies

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/zemartins81/devbookWebApp/src/config"
)

var s *securecookie.SecureCookie

// Configurar utiliza as variáveis de ambiente para o SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Grava as informações de login
func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	dadosCodificados, err := s.Encode("dados", dados)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
