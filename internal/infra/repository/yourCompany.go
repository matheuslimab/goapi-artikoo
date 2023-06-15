package repository

import (
	"database/sql"

	"github.com/matheuslimab/artikoo/api/src/models"
)

type dBCompany struct {
	db *sql.DB
}

func NewRepositoryYourCompany(db *sql.DB) *dBCompany {
	return &dBCompany{db}
}

func (repository *dBCompany) CreateCompany(company models.YourCompany) (string, error) {

	statament, err := repository.db.Prepare("INSERT INTO YourCompany (id_company, razao_social, nome_fantasia, cnpj, inscricao_estadual, endereco, telefone, email, site, contato_principal, segmento_de_atuacao, observacao, id_user, status_company) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer statament.Close()

	_, erro := statament.Exec(
		company.Id_company,
		company.RazaoSocial,
		company.NomeFantasia,
		company.Cnpj,
		company.InscricaoEstadual,
		company.Endereco,
		company.Telefone,
		company.Email,
		company.Site,
		company.ContatoPrincipal,
		company.SegmentoDeAtuacao,
		company.Observacao,
		company.IdUser,
		company.StatusCompany,
	)
	if erro != nil {
		return "", err
	}

	return company.Id_company, nil
}

func (repository *dBCompany) GetCompany(id string) (models.YourCompany, error) {
	rows, err := repository.db.Query("SELECT id_company, razao_social, nome_fantasia, cnpj, inscricao_estadual, endereco, telefone, email, site, contato_principal, segmento_de_atuacao, observacao, id_user, status_company FROM YourCompany WHERE id_company = ?", id)
	if err != nil {
		return models.YourCompany{}, err
	}
	defer rows.Close()

	var company models.YourCompany

	for rows.Next() {
		if err := rows.Scan(
			&company.Id_company,
			&company.RazaoSocial,
			&company.NomeFantasia,
			&company.Cnpj,
			&company.InscricaoEstadual,
			&company.Endereco,
			&company.Telefone,
			&company.Email,
			&company.Site,
			&company.ContatoPrincipal,
			&company.SegmentoDeAtuacao,
			&company.Observacao,
			&company.IdUser,
			&company.StatusCompany,
		); err != nil {
			return models.YourCompany{}, err
		}
	}

	return company, nil
}

func (repository *dBCompany) GetCompanies() ([]models.YourCompany, error) {

	rows, err := repository.db.Query("SELECT id_company, razao_social, nome_fantasia, cnpj, inscricao_estadual, endereco, telefone, email, site, contato_principal, segmento_de_atuacao, observacao, id_user, status_company FROM YourCompany")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []models.YourCompany

	for rows.Next() {
		var company models.YourCompany

		if err := rows.Scan(
			&company.Id_company,
			&company.RazaoSocial,
			&company.NomeFantasia,
			&company.Cnpj,
			&company.InscricaoEstadual,
			&company.Endereco,
			&company.Telefone,
			&company.Email,
			&company.Site,
			&company.ContatoPrincipal,
			&company.SegmentoDeAtuacao,
			&company.Observacao,
			&company.IdUser,
			&company.StatusCompany,
		); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (repository *dBCompany) UpdateCompany(company models.YourCompany, id string) error {
	statement, err := repository.db.Prepare("UPDATE YourCompany SET razao_social = ?, nome_fantasia = ?, cnpj = ?, inscricao_estadual = ?, endereco = ?, telefone = ?, email = ?, site = ?, contato_principal = ?, segmento_de_atuacao = ?, observacao = ?, id_user = ?, status_company = ? WHERE id_company = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(
		company.RazaoSocial,
		company.NomeFantasia,
		company.Cnpj,
		company.InscricaoEstadual,
		company.Endereco,
		company.Telefone,
		company.Email,
		company.Site,
		company.ContatoPrincipal,
		company.SegmentoDeAtuacao,
		company.Observacao,
		company.IdUser,
		company.StatusCompany,
		id,
	); err != nil {
		return err
	}

	return nil
}
