package models

import (
	"strings"

	"github.com/matheuslimab/artikoo/api/resources/consts"
)

type Clients struct {
	Id_client                string `json:"id_client"`
	Type_client              uint64 `json:"type_client"`
	Nome_razao_social_client string `json:"nome_razao_social_client"`

	Cpf_cnpj_client                     string `json:"cpf_cnpj_client"`
	Rg_ie_client                        string `json:"rg_ie_client"`
	Data_nascimento_dataFundacao_client string `json:"data_nascimento_dataFundacao_client"`

	Endereco_client string `json:"endereco_client"`
	Telefone_client string `json:"telefone_client"`
	Celular_client  string `json:"celular_client"`
	Email_client    string `json:"email_client"`

	Site_client                string `json:"site_client"`
	Contato_principal_client   string `json:"contato_principal_client"`
	Segmento_de_atuacao_client uint64 `json:"segmento_de_atuacao_client"`

	Limite_credito_client      string `json:"limite_credito_client"`
	Condicoes_pagamento_client string `json:"condicoes_pagamento_client"`
	Status_cliente_client      uint64 `json:"status_cliente_client"`
	Observacao_client          string `json:"observacao_client"`

	Created_at string `json:"created_at"`
	Update_at  string `json:"update_at"`

	EmpresaID_clients uint64 `json:"empresaID_clients`
	User_id           uint64 `json:"user_id`
}

func (t *Clients) Check() error {
	err := t.validar()
	if err != nil {
		return err
	}

	t.formatar()
	return nil
}

func (t *Clients) formatar() {
	t.Nome_razao_social_client = strings.TrimSpace(t.Nome_razao_social_client)
}

func (t *Clients) validar() error {

	if len(t.Nome_razao_social_client) > 100 {
		return consts.ErrMaximumLengthNameClient
	}

	if len(t.Nome_razao_social_client) <= 3 {
		return consts.ErrMinimumLengthNameClient
	}

	return nil

}
