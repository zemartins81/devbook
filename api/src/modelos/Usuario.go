package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário no sistema
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar valida e formata o usuário.
//
// Retorna um erro se a validação falhar.
// Retorna nil se a validação for bem-sucedida.
func (u *Usuario) Preparar(etapa string) error {
	if erro := u.formatar(etapa); erro != nil {
		return erro
	}

	if erro := u.validar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (u *Usuario) validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("O nome é obrigatório")
	}
	if u.Nick == "" {
		return errors.New("O nick é obrigatório")
	}
	if u.Email == "" {
		return errors.New("O Email é obrigatório")
	}
	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}
	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("A senha é obrigatória")
	}

	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(u.Senha)
		if erro != nil {
			return erro
		}

		u.Senha = string(senhaComHash)
	}

	return nil
}
