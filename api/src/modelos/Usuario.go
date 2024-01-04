package modelos

import (
	"errors"
	"strings"
	"time"
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
func (u *Usuario) Preparar() error {
	u.formatar()

	if erro := u.validar(); erro != nil {
		return erro
	}

	return nil
}

func (u *Usuario) validar() error {
	if u.Nome == "" {
		return errors.New("O nome é obrigatório")
	}
	if u.Nick == "" {
		return errors.New("O nick é obrigatório")
	}
	if u.Email == "" {
		return errors.New("O Email é obrigatório")
	}
	if u.Senha == "" {
		return errors.New("A senha é obrigatória")
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
