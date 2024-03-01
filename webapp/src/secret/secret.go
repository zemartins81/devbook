package secret

import "github.com/gorilla/securecookie"

func GerarSecretKey() []byte {
	secretKey := securecookie.GenerateRandomKey(16)
	return secretKey
}
