package rotas

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/entity"
)

var rotaYourCompany = []Rota{
	{
		URI:                "/createYourCompany",
		Metodo:             http.MethodPost,
		Funcao:             entity.CreateYourCompany,
		RequerAutenticacao: false,
	},
	{
		URI:                "/getYourCompany/{id}",
		Metodo:             http.MethodGet,
		Funcao:             entity.GetYourCompany,
		RequerAutenticacao: false,
	},
	{
		URI:                "/getAllYourCompany",
		Metodo:             http.MethodGet,
		Funcao:             entity.GetAllYourCompany,
		RequerAutenticacao: false,
	},
	{
		URI:                "/updateYourCompany/{id}",
		Metodo:             http.MethodPut,
		Funcao:             entity.UpdateYourCompany,
		RequerAutenticacao: false,
	},
}
