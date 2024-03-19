package modelos

import "time"

// Usuario representa uma pessoa usando o DevBook
type Usuario struct {
	ID          uint64       `json:"id,omitempty"`
	Nome        string       `json:"nome,omitempty"`
	Email       string       `json:"email,omitempty"`
	Nick        string       `json:"nick,omitempty"`
	CriadoEm    time.Time    `json:"criadoEm,omitempty"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}
