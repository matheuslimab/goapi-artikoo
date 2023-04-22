package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Data struct {
	token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	var usuario models.Usuario
	if err := json.Unmarshal(body, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)
	user_db, err := new_repository.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.Verify(user_db.Senha, usuario.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.CreateToken(user_db.ID)

	respostas.JSON(w, http.StatusOK, token)
}
