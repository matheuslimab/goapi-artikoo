package entity

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/config"
	"github.com/matheuslimab/artikoo/api/resources/consts"
)

func AuthUserAPI(r *http.Request) error {
	xAuthApiUser := r.Header.Get("X-AUTH-API-USER")
	xAuthApiPass := r.Header.Get("X-AUTH-API-PASS")
	if xAuthApiUser != "A123456" && xAuthApiPass != "88888888" {
		return consts.ErrAuthLoginAPI
	}

	return nil
}

func AuthorizeHeaderRequest(w http.ResponseWriter, r *http.Request) error {
	xAuthApiCrm := r.Header.Get("X-AUTH-API-CRM")

	if xAuthApiCrm != config.Api_key {
		return consts.ErrAuthApiKeyAPI
	}

	return nil
}
