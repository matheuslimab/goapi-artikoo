package models

import (
	"strings"

	"github.com/matheuslimab/artikoo/api/resources/consts"
)

type YourCompany struct {
	Id_company        string `json:"id_company"`
	RazaoSocial       string `json:"razao_social"`
	NomeFantasia      string `json:"nome_fantasia"`
	Cnpj              string `json:"cnpj"`
	InscricaoEstadual string `json:"inscricao_estadual"`
	Endereco          string `json:"endereco"`
	Telefone          string `json:"telefone"`
	Email             string `json:"email"`
	Site              string `json:"site"`
	ContatoPrincipal  string `json:"contato_principal"`
	SegmentoDeAtuacao string `json:"segmento_de_atuacao"`
	Observacao        string `json:"observacao"`
	IdUser            string `json:"id_user"`
	CreatedAt         string `json:"created_at"`
	UpdateAt          string `json:"update_at"`
	StatusCompany     uint64 `json:"status_company"`
}

func (t *YourCompany) Check() error {
	err := t.validar()
	if err != nil {
		return err
	}

	t.formatar()
	return nil
}

func (t *YourCompany) formatar() {
	t.RazaoSocial = strings.TrimSpace(t.RazaoSocial)
}

func (t *YourCompany) validar() error {
	if len(t.RazaoSocial) > consts.MaximumLengthRazaoSocialCompany {
		return consts.ErrMaximumLengthrRazaoSocialCompany
	}

	if len(t.RazaoSocial) <= consts.MinimumLengthRazaoSocialCompany {
		return consts.ErrMinimumLengthRazaoSocialCompany
	}

	return nil
}
