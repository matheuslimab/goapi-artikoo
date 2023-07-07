package rotas

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/entity"
)

var rotaPersonalData = []Rota{
	{
		URI:                "/createPersonalData",
		Metodo:             http.MethodPost,
		Funcao:             entity.CreatePersonalData,
		RequerAutenticacao: false,
	},
	{
		URI:                "/getPersonalData/{id}",
		Metodo:             http.MethodGet,
		Funcao:             entity.GetPersonalData,
		RequerAutenticacao: false,
	},
}
