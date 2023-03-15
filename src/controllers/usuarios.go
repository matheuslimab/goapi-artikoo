package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Essa funcao é responsavel por criar um usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario

	if err := json.Unmarshal(body, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("cadastro"); err != nil {
		fmt.Println("Error line #33 file: #src/controllers/usuarios.go")
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)
	usuariosQueryDB, err := new_repository.VerifyMail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if usuariosQueryDB.Email == usuario.Email {
		fmt.Println("O e-mail já existe cadastrado.")
		respostas.JSON(w, http.StatusConflict, true)
	} else {
		usuarioId, err := new_repository.Criar(usuario)
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
			return
		}

		usuario.ID = usuarioId

		respostas.JSON(w, http.StatusCreated, usuario)
	}

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := database.Connect()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	new_repository := repository.NewRepository(db)
	usuarios, err := new_repository.Buscar(nomeOuNick)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametro["id"], 10, 64)
	if err != nil {
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
	usuario, err := new_repository.BuscarPorID(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametro["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	usuarioIDFromToken, err := auth.GetUIDFromToken(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIDFromToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um id que não é seu"))
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err := json.Unmarshal(body, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := usuario.Preparar("atualizar"); err != nil {
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
	err = new_repository.AtualizarUsuario(usuarioID, usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIDFromToken, err := auth.GetUIDFromToken(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIDFromToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um id que não é seu"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)

	err = new_repository.DeletarUsuario(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
