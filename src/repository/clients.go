package repository

import (
	"database/sql"

	"github.com/matheuslimab/artikoo/api/resources/consts"
	"github.com/matheuslimab/artikoo/api/src/models"
)

type dB struct {
	db *sql.DB
}

func NewRepositoryClient(db *sql.DB) *dB {
	return &dB{db}
}

func (repository *dB) GetAllClients() ([]models.Clients, error) {

	row, err := repository.db.Query(`SELECT id_client, type_client, nome_razao_social_client, cpf_cnpj_client, rg_ie_client, data_nascimento_dataFundacao_client, endereco_client, telefone_client, celular_client, email_client, site_client, contato_principal_client, segmento_de_atuacao_client, limite_credito_client, condicoes_pagamento_client, status_cliente_client, observacao_client, created_at, update_at, empresaID_clients, user_id FROM Clients`)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var clients []models.Clients

	for row.Next() {

		var client models.Clients

		err := row.Scan(
			&client.Id_client,
			&client.Type_client,
			&client.Nome_razao_social_client,
			&client.Cpf_cnpj_client,
			&client.Rg_ie_client,
			&client.Data_nascimento_dataFundacao_client,
			&client.Endereco_client,
			&client.Telefone_client,
			&client.Celular_client,
			&client.Email_client,
			&client.Site_client,
			&client.Contato_principal_client,
			&client.Segmento_de_atuacao_client,
			&client.Limite_credito_client,
			&client.Condicoes_pagamento_client,
			&client.Status_cliente_client,
			&client.Observacao_client,
			&client.Created_at,
			&client.Update_at,
			&client.EmpresaID_clients,
			&client.User_id,
		)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)

	}

	return clients, nil

}

func (repository *dB) GetClient(clientID string) (models.Clients, error) {
	row, err := repository.db.Query(`SELECT id_client, type_client, nome_razao_social_client, cpf_cnpj_client, rg_ie_client, data_nascimento_dataFundacao_client, endereco_client, telefone_client, celular_client, email_client, site_client, contato_principal_client, segmento_de_atuacao_client, limite_credito_client, condicoes_pagamento_client, status_cliente_client, observacao_client, created_at, update_at, empresaID_clients, user_id FROM Clients WHERE id_client = ?`, clientID)
	if err != nil {
		return models.Clients{}, err
	}
	defer row.Close()

	var clients models.Clients

	if row.Next() {
		err := row.Scan(
			&clients.Id_client,
			&clients.Type_client,
			&clients.Nome_razao_social_client,
			&clients.Cpf_cnpj_client,
			&clients.Rg_ie_client,
			&clients.Data_nascimento_dataFundacao_client,
			&clients.Endereco_client,
			&clients.Telefone_client,
			&clients.Celular_client,
			&clients.Email_client,
			&clients.Site_client,
			&clients.Contato_principal_client,
			&clients.Segmento_de_atuacao_client,
			&clients.Limite_credito_client,
			&clients.Condicoes_pagamento_client,
			&clients.Status_cliente_client,
			&clients.Observacao_client,
			&clients.Created_at,
			&clients.Update_at,
			&clients.EmpresaID_clients,
			&clients.User_id,
		)
		if err != nil {
			return models.Clients{}, err
		}
	}

	return clients, nil
}

func (repository *dB) CreateClient(client models.Clients) (string, error) {
	statament, err := repository.db.Prepare("INSERT INTO Clients ( id_client, type_client, nome_razao_social_client, cpf_cnpj_client, rg_ie_client, data_nascimento_dataFundacao_client, endereco_client, telefone_client, celular_client, email_client, site_client, contato_principal_client, segmento_de_atuacao_client, limite_credito_client, condicoes_pagamento_client, status_cliente_client, observacao_client, empresaID_clients, user_id) VALUES (?, ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer statament.Close()

	_, erro := statament.Exec(
		client.Id_client,
		client.Type_client,
		client.Nome_razao_social_client,
		client.Cpf_cnpj_client,
		client.Rg_ie_client,
		client.Data_nascimento_dataFundacao_client,
		client.Endereco_client,
		client.Telefone_client,
		client.Celular_client,
		client.Email_client,
		client.Site_client,
		client.Contato_principal_client,
		client.Segmento_de_atuacao_client,
		client.Limite_credito_client,
		client.Condicoes_pagamento_client,
		client.Status_cliente_client,
		client.Observacao_client,
		client.EmpresaID_clients,
		client.User_id,
	)
	if erro != nil {
		return "", err
	}

	return client.Id_client, nil
}

func (repository *dB) UpdateClient(client models.Clients, uid_client string) error {

	clientsResponse, erro := repository.GetClient(uid_client)
	if erro != nil {
		return erro
	}

	if clientsResponse.Id_client != "" {

		statament, err := repository.db.Prepare("UPDATE Clients SET type_client = ?, nome_razao_social_client = ?, cpf_cnpj_client = ?, rg_ie_client = ?, data_nascimento_dataFundacao_client = ?, endereco_client = ?, telefone_client = ?, celular_client = ?, email_client = ?, site_client = ?, contato_principal_client = ?, segmento_de_atuacao_client = ?, limite_credito_client = ?, condicoes_pagamento_client = ?, status_cliente_client = ?, observacao_client = ? WHERE id_client = ?")
		if err != nil {
			return err
		}
		defer statament.Close()

		if err = client.Check(); err != nil {
			return err
		}

		_, err = statament.Exec(
			client.Type_client,
			client.Nome_razao_social_client,
			client.Cpf_cnpj_client,
			client.Rg_ie_client,
			client.Data_nascimento_dataFundacao_client,
			client.Endereco_client,
			client.Telefone_client,
			client.Celular_client,
			client.Email_client,
			client.Site_client,
			client.Contato_principal_client,
			client.Segmento_de_atuacao_client,
			client.Limite_credito_client,
			client.Condicoes_pagamento_client,
			client.Status_cliente_client,
			client.Observacao_client,
			uid_client,
		)

		if err != nil {
			return err
		}

		return nil

	} else {
		return consts.ErrNotFoundClient
	}
}
