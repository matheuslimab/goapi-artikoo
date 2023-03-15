package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id, omitempty"`
	Nome     string    `nome:"nome, omitempty"`
	Nick     string    `nick:"nick, omitempty"`
	Email    string    `email:"email, omitempty"`
	Senha    string    `senha:"senha, omitempty"`
	CriadoEm time.Time `criadoEm:"criadoEm, omitempty"`
}

func (u *Usuario) Preparar(etapa string) error {
	if erro := u.validar(etapa); erro != nil {
		return erro
	}

	err := u.formatar(etapa)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usuario) validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("o campo nome nao pode ficar em branco")
	} else if u.Nick == "" {
		return errors.New("o campo Nick é obrigatório")
	}

	if u.Email == "" {
		return errors.New("o campo Email é obrigatório")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("o e-mail inserido esta incorreto")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("o campo Senha é obrigatório")
	}

	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Senha = strings.TrimSpace(u.Senha)

	if etapa == "cadastro" {
		senhaHashed, erro := security.Hash(u.Senha)
		if erro != nil {
			return erro
		}

		u.Senha = string(senhaHashed)
	}

	return nil
}
