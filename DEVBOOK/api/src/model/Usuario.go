package model

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json: "id,omitempty"`
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "senha, omitempty"`
	CriadoEm time.Time `json: "CriadoEm, omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido.
func (u *Usuario) Preparar() error {
	if erro := u.validar(); erro != nil {
		return erro
	}

	u.formatar()
	return nil
}

func (u *Usuario) validar() error {
	if u.Nome == "" || u.Nick == "" || u.Email == "" || u.Senha == "" {
		return errors.New("Os campos são obrigatórios.")
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
