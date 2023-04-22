package models

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id, omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autor_id,omitempty"`
	AutorNick string    `json:"autor_nick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

func (pub *Publicacao) Preparar() error {
	if erro := pub.validar(); erro != nil {
		return erro
	}

	pub.Formatar()
	return nil
}

func (pub *Publicacao) Formatar() {
	pub.Titulo = strings.TrimSpace(pub.Titulo)
	pub.Conteudo = strings.TrimSpace(pub.Conteudo)
}

func (pub *Publicacao) validar() error {

	if len(pub.Titulo) <= 1 {
		return errors.New("você precisa escrever mais, pois 1(um) caractere nao sera aceito")
	}

	if len(pub.Conteudo) <= 1 {
		return errors.New("voce precisa escrever mais, pois 1(um) caractere nao sera aceito")
	}

	if pub.Titulo == "" {
		return errors.New("o titulo e obrigatorio e nao pode estar em branco")
	}

	if pub.Conteudo == "" {
		return errors.New("o conteudo e obrigatorio e nao pode estar em branco")
	}

	return nil
}
