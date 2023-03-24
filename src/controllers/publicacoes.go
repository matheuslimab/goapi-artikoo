package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/helpers"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	uid, err := auth.GetUIDFromToken(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusBadGateway, err)
		return
	}

	var publicacao models.Publicacao

	publicacao.AutorID = uid

	if err := json.Unmarshal(body, &publicacao); err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := publicacao.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	new_repository := repository.NewRepositoryPublicacoes(db)
	pubID, err := new_repository.Criar(publicacao)
	if err != nil {
		helpers.Err(w, http.StatusInternalServerError, err)
	}
	publicacao.ID = pubID

	respostas.JSON(w, http.StatusCreated, publicacao)

}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
