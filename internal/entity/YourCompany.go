package entity

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheuslimab/artikoo/api/internal/infra/repository"
	pkgEntity "github.com/matheuslimab/artikoo/api/pkg/entity"
	"github.com/matheuslimab/artikoo/api/src/database"
	"github.com/matheuslimab/artikoo/api/src/helpers"
	"github.com/matheuslimab/artikoo/api/src/models"
)

func GetAllYourCompany(w http.ResponseWriter, r *http.Request) {
	err := pkgEntity.AuthUserAPI(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	err = pkgEntity.AuthorizeHeaderRequest(w, r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	new_repository := repository.NewRepositoryYourCompany(db)
	companies, err := new_repository.GetCompanies()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, companies)
}

func GetYourCompany(w http.ResponseWriter, r *http.Request) {
	err := pkgEntity.AuthUserAPI(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	err = pkgEntity.AuthorizeHeaderRequest(w, r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	new_repository := repository.NewRepositoryYourCompany(db)
	company, err := new_repository.GetCompany(parametros["id"])
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, company)
}

func CreateYourCompany(w http.ResponseWriter, r *http.Request) {

	err := pkgEntity.AuthUserAPI(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	err = pkgEntity.AuthorizeHeaderRequest(w, r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	var yourCompany models.YourCompany

	if err := json.Unmarshal(body, &yourCompany); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := yourCompany.Check(); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	new_repository := repository.NewRepositoryYourCompany(db)
	result, err := new_repository.SelectCompanyByIdUser(yourCompany.IdUser)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if result.Id_company != "" {
		err := new_repository.UpdateCompany(yourCompany, result.Id_company)
		if err != nil {
			pkgEntity.Erro(w, http.StatusBadRequest, err)
			return
		}

		pkgEntity.JSON(w, http.StatusOK, nil)
	} else {
		yourCompany.Id_company = pkgEntity.GenerateNewID()
		yourCompanyID, err := new_repository.CreateCompany(yourCompany)
		if err != nil {
			pkgEntity.Erro(w, http.StatusBadRequest, err)
			return
		}

		pkgEntity.JSON(w, http.StatusCreated, yourCompanyID)
	}

}

func UpdateYourCompany(w http.ResponseWriter, r *http.Request) {
	err := pkgEntity.AuthUserAPI(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	err = pkgEntity.AuthorizeHeaderRequest(w, r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	var company models.YourCompany
	if err := json.Unmarshal(body, &company); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := company.Check(); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	new_repository := repository.NewRepositoryYourCompany(db)
	err = new_repository.UpdateCompany(company, parametros["id"])
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	pkgEntity.JSON(w, http.StatusCreated, company)
}
