package secret

import (
	"crypto/rand"
	"encoding/base64"
)

func GeraSecret() (string, error) {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		return "", erro
	}

	// Converte a chave para uma string base64.
	chaveBase64 := base64.StdEncoding.EncodeToString(chave)
	return chaveBase64, nil
}
