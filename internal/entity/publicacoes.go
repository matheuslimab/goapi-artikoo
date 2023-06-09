package entity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	pkgEntity "github.com/matheuslimab/artikoo/api/pkg/entity"
	"github.com/matheuslimab/artikoo/api/src/auth"
	"github.com/matheuslimab/artikoo/api/src/database"
	"github.com/matheuslimab/artikoo/api/src/helpers"
	"github.com/matheuslimab/artikoo/api/src/models"
	"github.com/matheuslimab/artikoo/api/src/repository"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	uid, err := auth.GetUIDFromToken(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadGateway, err)
		return
	}

	var publicacao models.Publicacao

	publicacao.AutorID = uid

	if err := json.Unmarshal(body, &publicacao); err != nil {
		pkgEntity.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := publicacao.Preparar(); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
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

	pkgEntity.JSON(w, http.StatusCreated, publicacao)

}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.GetUIDFromToken(r)
	if err != nil {
		pkgEntity.JSON(w, http.StatusUnauthorized, err)
	}

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	new_repository := repository.NewRepositoryPublicacoes(db)
	publicidades, err := new_repository.Buscar(uid)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, publicidades)
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	helpers.Err(w, http.StatusInternalServerError, err)
	defer db.Close()

	new_repository := repository.NewRepositoryPublicacoes(db)
	publicacao, erro := new_repository.BuscarPorID(publicacaoID)
	if erro != nil {
		helpers.Err(w, http.StatusInternalServerError, err)
	}

	pkgEntity.JSON(w, http.StatusOK, publicacao)
}

// AtualizarPublicacao altera os dados de uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.GetUIDFromToken(r)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewRepositoryPublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		pkgEntity.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacao.Preparar(); erro != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(publicacaoID, publicacao); erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)
}

// DeletarPublicacao exclui os dados de uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.GetUIDFromToken(r)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewRepositoryPublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		pkgEntity.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar uma publicação que não seja sua"))
		return
	}

	if erro = repositorio.Deletar(publicacaoID); erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)
}

// BuscarPublicacoesPorUsuario traz todas as publicações de um usuário específico
func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewRepositoryPublicacoes(db)
	publicacoes, erro := repositorio.BuscarPorUsuario(usuarioID)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, publicacoes)
}

// CurtirPublicacao adiciona uma curtida na publicação
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewRepositoryPublicacoes(db)
	if erro = repositorio.Curtir(publicacaoID); erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)
}

// DescurtirPublicacao subtrai uma curtida na publicação
func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewRepositoryPublicacoes(db)
	if erro = repositorio.Descurtir(publicacaoID); erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)
}
