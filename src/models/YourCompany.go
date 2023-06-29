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
	t.NomeFantasia = strings.TrimSpace(t.NomeFantasia)
	t.Cnpj = strings.TrimSpace(t.Cnpj)
	t.InscricaoEstadual = strings.TrimSpace(t.InscricaoEstadual)
	t.Endereco = strings.TrimSpace(t.Endereco)
	t.Telefone = strings.TrimSpace(t.Telefone)
	t.Email = strings.TrimSpace(t.Email)
	t.Site = strings.TrimSpace(t.Site)
	t.ContatoPrincipal = strings.TrimSpace(t.ContatoPrincipal)
	t.SegmentoDeAtuacao = strings.TrimSpace(t.SegmentoDeAtuacao)
	t.Observacao = strings.TrimSpace(t.Observacao)
}

func (t *YourCompany) validar() error {

	if len(t.RazaoSocial) > consts.MaximumLengthRazaoSocialCompany {
		return consts.ErrMaximumLengthrRazaoSocialCompany
	}

	if len(t.RazaoSocial) <= consts.MinimumLengthRazaoSocialCompany {
		return consts.ErrMinimumLengthRazaoSocialCompany
	}

	if len(t.NomeFantasia) > consts.MaximumLengthNomeFantasiaCompany {
		return consts.ErrMaximumLengthrNomeFantasiaCompany
	}

	if len(t.NomeFantasia) <= consts.MinimumLengthNomeFantasiaCompany {
		return consts.ErrMinimumLengthNomeFantasiaCompany
	}

	if len(t.Cnpj) > consts.MaximumLengthCnpjCompany {
		return consts.ErrMaximumLengthCnpjCompany
	}

	if len(t.Cnpj) <= consts.MinimumLengthCnpjCompany {
		return consts.ErrMinimumLengthCnpjCompany
	}

	if len(t.InscricaoEstadual) > consts.MaximumLengthIeCompany {
		return consts.ErrMaximumLengthIeCompany
	}

	if len(t.InscricaoEstadual) <= consts.MinimumLengthIeCompany {
		return consts.ErrMinimumLengthIeCompany
	}

	if len(t.Endereco) > consts.MaximumLengthAddressCompany {
		return consts.ErrMaximumLengthAddressCompany
	}

	if len(t.Endereco) <= consts.MinimumLengthAddressCompany {
		return consts.ErrMinimumLengthAddressCompany
	}

	if len(t.Telefone) > consts.MaximumLengthPhoneCompany {
		return consts.ErrMaximumLengthPhoneCompany
	}

	if len(t.Telefone) <= consts.MinimumLengthPhoneCompany {
		return consts.ErrMinimumLengthPhoneCompany
	}

	if len(t.Email) > consts.MaximumLengthEmailCompany {
		return consts.ErrMaximumLengthEmailCompany
	}

	if len(t.Email) <= consts.MinimumLengthEmailCompany {
		return consts.ErrMinimumLengthEmailCompany
	}

	if len(t.Site) > consts.MaximumLengthSiteCompany {
		return consts.ErrMaximumLengthSiteCompany
	}

	if len(t.Site) <= consts.MinimumLengthSiteCompany {
		return consts.ErrMinimumLengthSiteCompany
	}

	if len(t.ContatoPrincipal) > consts.MaximumLengthContactCompany {
		return consts.ErrMaximumLengthContactCompany
	}

	if len(t.ContatoPrincipal) <= consts.MinimumLengthContactCompany {
		return consts.ErrMinimumLengthContactCompany
	}

	if len(t.SegmentoDeAtuacao) > consts.MaximumLengthSegmentCompany {
		return consts.ErrMaximumLengthSegmentCompany
	}

	if len(t.SegmentoDeAtuacao) <= consts.MinimumLengthSegmentCompany {
		return consts.ErrMinimumLengthSegmentCompany
	}

	if len(t.Observacao) > consts.MaximumLengthObservationCompany {
		return consts.ErrMaximumLengthObservationCompany
	}

	if len(t.Observacao) <= consts.MinimumLengthObservationCompany {
		return consts.ErrMinimumLengthObservationCompany
	}

	return nil
}
