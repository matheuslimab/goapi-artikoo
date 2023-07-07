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

func CreatePersonalData(w http.ResponseWriter, r *http.Request) {
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

	var personalData models.PersonalData

	if err := json.Unmarshal(body, &personalData); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := personalData.CheckData(); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	new_repository := repository.NewRepositoryPersonalData(db)
	result, err := new_repository.SelectCompanyByIdUser(personalData.IdUser)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if result.IdUser != "" {
		err := new_repository.UpdatePersonalData(personalData, result.IdUser)
		if err != nil {
			pkgEntity.Erro(w, http.StatusInternalServerError, err)
			return
		}

		pkgEntity.JSON(w, http.StatusOK, nil)
	} else {
		personalData.IdPersonalData = pkgEntity.GenerateNewID()
		response, err := new_repository.CreatePersonalData(personalData)
		if err != nil {
			pkgEntity.Erro(w, http.StatusBadRequest, err)
			return
		}

		pkgEntity.JSON(w, http.StatusCreated, response)
	}
}

func UpdatePersonalData(w http.ResponseWriter, r *http.Request) {
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

	new_repository := repository.NewRepositoryPersonalData(db)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	var personalData models.PersonalData

	if err := json.Unmarshal(body, &personalData); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := personalData.CheckData(); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	err = new_repository.UpdatePersonalData(personalData, parametros["id"])
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, nil)
}

func GetPersonalData(w http.ResponseWriter, r *http.Request) {
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

	new_repository := repository.NewRepositoryPersonalData(db)

	result, err := new_repository.GetPersonalData(parametros["id"])
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, result)
}
