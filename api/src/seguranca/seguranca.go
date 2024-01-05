package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash gera um hash a partir de uma senha usando bcrypt.
// Retorna o hash como um slice de bytes.
func Hash(senha string) ([]byte, error) {
	// Gera o hash da senha utilizando o custo padrão.
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// VerificarSenha compara uma senha com hash com uma senha em texto puro.
// Retorna um erro se as senhas não coincidirem.
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
