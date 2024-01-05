package secret

import (
	"crypto/rand"
)

func GeraSecret() ([]byte, error) {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		return nil, erro
	}

	return chave, nil
}
