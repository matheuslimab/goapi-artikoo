package entity

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/infra/repository"
	pkgEntity "github.com/matheuslimab/artikoo/api/pkg/entity"
	"github.com/matheuslimab/artikoo/api/src/database"
	"github.com/matheuslimab/artikoo/api/src/helpers"
)

func GetAllSegments(w http.ResponseWriter, r *http.Request) {
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

	new_repository := repository.NewRepositorySegments(db)
	clients, err := new_repository.GetAllSegments()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pkgEntity.JSON(w, http.StatusOK, clients)
}
