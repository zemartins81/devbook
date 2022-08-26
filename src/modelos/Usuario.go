package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário do sistema
type Usuario struct {
	ID       uint64    `json:"id, omitempty"`
	Nome     string    `json:"nome, omitempty"`
	Nick     string    `json:"nick, omitempty"`
	Email    string    `json:"email, omitempty"`
	Senha    string    `json:"senha, omitempty"`
	CriadoEm time.Time `json:"CriadoEm, omitempty`
}

// Preparar vai chamar os metodos para validar e formatar o usuario
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O nome não pode ficar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick não pode ficar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O email não pode ficar em branco")
	}
	if usuario.Senha == "" {
		return errors.New("A senha não pode ficar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
