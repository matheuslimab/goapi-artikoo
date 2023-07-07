package repository

import (
	"database/sql"

	"github.com/matheuslimab/artikoo/api/src/models"
)

type dbRepositoryPersonalData struct {
	db *sql.DB
}

func NewRepositoryPersonalData(db *sql.DB) *dbRepositoryPersonalData {
	return &dbRepositoryPersonalData{db}
}

func (repository *dbRepositoryPersonalData) SelectCompanyByIdUser(idUser string) (models.PersonalData, error) {
	rows, err := repository.db.Query("SELECT idPersonalData_dataUser, name_dataUser, email_dataUser, cpf_dataUser, rg_dataUser, dataNascimento_dataUser, cargo_dataUser, timeDeparture_DataUser, aboutMe_dataUser, ausencia_dataUser, genero_dataUser, endereco_dataUser, idUser_dataUser FROM PersonalData WHERE idUser_dataUser = ?", idUser)
	if err != nil {
		return models.PersonalData{}, err
	}
	defer rows.Close()

	var personalData models.PersonalData

	for rows.Next() {
		if err := rows.Scan(
			&personalData.IdPersonalData,
			&personalData.Name,
			&personalData.Email,
			&personalData.Cpf,
			&personalData.Rg,
			&personalData.DataNascimento,
			&personalData.Cargo,
			&personalData.TimeDeparture,
			&personalData.AboutMe,
			&personalData.Ausencia,
			&personalData.Genero,
			&personalData.EnderecoPessoal,
			&personalData.IdUser,
		); err != nil {
			return models.PersonalData{}, err
		}
	}

	return personalData, nil
}

func (repository *dbRepositoryPersonalData) CreatePersonalData(personalData models.PersonalData) (string, error) {

	statament, err := repository.db.Prepare("INSERT INTO PersonalData (idPersonalData_dataUser, name_dataUser, email_dataUser, cpf_dataUser, rg_dataUser, dataNascimento_dataUser, cargo_dataUser, timeDeparture_DataUser, aboutMe_dataUser, ausencia_dataUser, genero_dataUser, endereco_dataUser, idUser_dataUser) VALUES(?, ?, ?,?, ?, ?, ?, ?, ?, ?,?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer statament.Close()

	_, erro := statament.Exec(
		personalData.IdPersonalData,
		personalData.Name,
		personalData.Email,
		personalData.Cpf,
		personalData.Rg,
		personalData.DataNascimento,
		personalData.Cargo,
		personalData.TimeDeparture,
		personalData.AboutMe,
		personalData.Ausencia,
		personalData.Genero,
		personalData.EnderecoPessoal,
		personalData.IdUser,
	)
	if erro != nil {
		return "", err
	}

	return personalData.IdPersonalData, nil
}

func (repository *dbRepositoryPersonalData) GetPersonalData(id string) (models.PersonalData, error) {

	var personalData models.PersonalData

	rows, err := repository.db.Query("SELECT * FROM PersonalData WHERE idUser_dataUser = ?", id)
	if err != nil {
		return personalData, err
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(
			&personalData.IdPersonalData,
			&personalData.Name,
			&personalData.Email,
			&personalData.Cpf,
			&personalData.Rg,
			&personalData.DataNascimento,
			&personalData.Cargo,
			&personalData.TimeDeparture,
			&personalData.AboutMe,
			&personalData.Ausencia,
			&personalData.Genero,
			&personalData.EnderecoPessoal,
			&personalData.IdUser,
			&personalData.Created_At,
			&personalData.Updated_At,
		); err != nil {
			return personalData, err
		}
	}

	return personalData, nil
}

func (repository *dbRepositoryPersonalData) UpdatePersonalData(personalData models.PersonalData, idUser string) error {

	statament, err := repository.db.Prepare("UPDATE PersonalData SET name_dataUser = ?, email_dataUser = ?, cpf_dataUser = ?, rg_dataUser = ?, dataNascimento_dataUser = ?, cargo_dataUser = ?, timeDeparture_DataUser = ?, aboutMe_dataUser = ?, ausencia_dataUser = ?, genero_dataUser = ?, endereco_dataUser = ?, idUser_dataUser = ?, updated_at_dataUser = ? WHERE idUser_dataUser = ?")
	if err != nil {
		return err
	}
	defer statament.Close()

	_, erro := statament.Exec(
		personalData.Name,
		personalData.Email,
		personalData.Cpf,
		personalData.Rg,
		personalData.DataNascimento,
		personalData.Cargo,
		personalData.TimeDeparture,
		personalData.AboutMe,
		personalData.Ausencia,
		personalData.Genero,
		personalData.EnderecoPessoal,
		personalData.IdUser,
		personalData.Updated_At,
		idUser,
	)
	if erro != nil {
		return err
	}

	return nil
}
