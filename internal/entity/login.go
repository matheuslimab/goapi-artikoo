package entity

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/infra/repository"
	pkgEntity "github.com/matheuslimab/artikoo/api/pkg/entity"
	"github.com/matheuslimab/artikoo/api/src/auth"
	"github.com/matheuslimab/artikoo/api/src/database"
	"github.com/matheuslimab/artikoo/api/src/models"
	"github.com/matheuslimab/artikoo/api/src/security"
)

type dataResponse struct {
	token string `json:"token"`
	models.Usuario
}

func Login(w http.ResponseWriter, r *http.Request) {

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

	var usuario models.Usuario
	if err := json.Unmarshal(body, &usuario); err != nil {
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
	user_db, err := new_repository.BuscarPorEmail(usuario.Email)
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.Verify(user_db.Senha, usuario.Senha); err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.CreateToken(user_db.ID)
	userInfos, err := new_repository.BuscarPorID(user_db.ID)
	if err != nil {
		pkgEntity.Erro(w, http.StatusUnauthorized, err)
		return
	}

	dataResponses := dataResponse{
		token,
		userInfos,
	}

	pkgEntity.JSON(w, http.StatusOK, dataResponses)
}
