package entity

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	pkgEntity "github.com/matheuslimab/artikoo/api/pkg/entity"
	"github.com/matheuslimab/artikoo/api/src/database"
	"github.com/matheuslimab/artikoo/api/src/helpers"
	"github.com/matheuslimab/artikoo/api/src/models"
	"github.com/matheuslimab/artikoo/api/src/repository"

	"github.com/gorilla/mux"
)

func GetAllClients(w http.ResponseWriter, r *http.Request) {
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

	new_repository := repository.NewRepositoryClient(db)
	clients, err := new_repository.GetAllClients()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, clients)
}

func GetClient(w http.ResponseWriter, r *http.Request) {

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
	// clientID, err := strconv.ParseInt(parametros["id"], 10, 64)
	// if err != nil {
	// 	pkgEntity.Erro(w, http.StatusBadRequest, err)
	// 	return
	// }

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	new_repository := repository.NewRepositoryClient(db)
	clients, err := new_repository.GetClient(parametros["id"])
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, clients)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {

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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	var clients models.Clients

	if err := json.Unmarshal(body, &clients); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := clients.Check(); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	clients.Id_client = pkgEntity.GenerateNewID()

	new_repository := repository.NewRepositoryClient(db)
	client_id, err := new_repository.CreateClient(clients)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	pkgEntity.JSON(w, http.StatusCreated, client_id)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {

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
	// uid_client, err := strconv.ParseInt(parametros["id"], 10, 64)
	// if err != nil {
	// 	pkgEntity.Erro(w, http.StatusBadRequest, err)
	// 	return
	// }

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	var clients models.Clients
	if err := json.Unmarshal(body, &clients); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := clients.Check(); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	new_repository := repository.NewRepositoryClient(db)
	err = new_repository.UpdateClient(clients, parametros["id"])
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	pkgEntity.JSON(w, http.StatusCreated, clients)
}
