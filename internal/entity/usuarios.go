package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	pkgEntity "github.com/matheuslimab/artikoo/api/pkg/entity"

	"github.com/matheuslimab/artikoo/api/src/auth"
	"github.com/matheuslimab/artikoo/api/src/database"
	"github.com/matheuslimab/artikoo/api/src/models"
	"github.com/matheuslimab/artikoo/api/src/repository"
	"github.com/matheuslimab/artikoo/api/src/security"

	"github.com/gorilla/mux"
)

// Essa funcao é responsavel por criar um usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario

	if err := json.Unmarshal(body, &usuario); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("cadastro"); err != nil {
		fmt.Println("Error line #33 file: #src/controllers/usuarios.go")
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)
	usuariosQueryDB, err := new_repository.VerifyMail(usuario.Email)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if usuariosQueryDB.Email == usuario.Email {
		fmt.Println("O e-mail já existe cadastrado.")
		pkgEntity.JSON(w, http.StatusConflict, true)
	} else {
		usuarioId, err := new_repository.Criar(usuario)
		if err != nil {
			pkgEntity.Erro(w, http.StatusInternalServerError, err)
			return
		}

		usuario.ID = usuarioId

		pkgEntity.JSON(w, http.StatusCreated, usuario)
	}

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	new_repository := repository.NewRepository(db)
	usuarios, err := new_repository.Buscar(nomeOuNick)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, usuarios)

}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametro["id"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)
	usuario, err := new_repository.BuscarPorID(usuarioID)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametro["id"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	usuarioIDFromToken, err := auth.GetUIDFromToken(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIDFromToken {
		pkgEntity.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um id que não é seu"))
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err := json.Unmarshal(body, &usuario); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := usuario.Preparar("atualizar"); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)
	err = new_repository.AtualizarUsuario(usuarioID, usuario)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIDFromToken, err := auth.GetUIDFromToken(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIDFromToken {
		pkgEntity.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um id que não é seu"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)

	err = new_repository.DeletarUsuario(usuarioID)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)
}

// SeguirUsuario permite que um usuario siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, err := auth.GetUIDFromToken(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorID == usuarioID {
		pkgEntity.Erro(w, http.StatusForbidden, errors.New("você não pode seguir você mesmo"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)
	followVerify, err := new_repository.VerifyFollowing(usuarioID, seguidorID)
	if err != nil {
		pkgEntity.Erro(w, http.StatusForbidden, err)
		return
	}

	// verificando se os dois usuario já se seguem, caso não se sigam o processo continua
	if followVerify.Usuario_ID == usuarioID && followVerify.Seguidor_ID == seguidorID {
		pkgEntity.Erro(w, http.StatusForbidden, errors.New("vocês já se seguem"))
		return
	}

	if err := new_repository.Seguir(usuarioID, seguidorID); err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)
}

func UnFollow(w http.ResponseWriter, r *http.Request) {

	seguidorID, err := auth.GetUIDFromToken(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorID == usuarioID {
		pkgEntity.Erro(w, http.StatusForbidden, errors.New("você não segue você mesmo, por tanto não é possível parar de seguir você mesmo"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepository(db)
	if erro := new_repository.UnFollow(usuarioID, seguidorID); erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	db.Close()

	new_repository := repository.NewRepository(db)
	seguidores, err := new_repository.BuscarSeguidores(usuarioID)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, seguidores)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	db.Close()

	new_repository := repository.NewRepository(db)
	usuarios, err := new_repository.BuscarSeguindo(usuarioID)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, usuarios)
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.GetUIDFromToken(r)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	if uid != usuarioID {
		pkgEntity.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não é seu"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	var senha models.Senha
	if erro := json.Unmarshal(body, &senha); erro != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	db.Close()

	new_repository := repository.NewRepository(db)
	passwordSavedDB, err := new_repository.BuscarSenha(usuarioID)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.Verify(passwordSavedDB, senha.Atual); err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, errors.New("a senha atual nao condiz com a que esta salva no banco"))
		return
	}

	senhaHashed, err := security.Hash(senha.Nova)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := new_repository.AtualizarSenha(usuarioID, string(senhaHashed)); err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusNoContent, nil)
}
