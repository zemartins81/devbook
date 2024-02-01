package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id, omitempty"`
	Titulo    string    `json:"titulo, omitempty"`
	Conteudo  string    `json:"conteudo, omitempty"`
	AutorID   uint64    `json:"autorId, omitempty"`
	AutorNick string    `json:"autorNick, omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm, omitempty"`
}

func (p *Publicacao) Preparar() error {
	p.formatar()

	if erro := p.validar(); erro != nil {
		return erro
	}

	return nil
}

func (p *Publicacao) validar() error {
	if p.Titulo == "" {
		return errors.New("O título é obrigadtório")
	}

	if p.Conteudo == "" {
		return errors.New("O conteúdo não pode ficar em branco")
	}

	return nil
}

func (p *Publicacao) formatar() {
	p.Titulo = strings.TrimSpace(p.Titulo)
	p.Conteudo = strings.TrimSpace(p.Conteudo)
}
