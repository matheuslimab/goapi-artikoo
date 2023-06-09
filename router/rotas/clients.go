package rotas

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/entity"
)

var rotaClients = []Rota{
	{
		URI:                "/createClient",
		Metodo:             http.MethodPost,
		Funcao:             entity.CreateClient,
		RequerAutenticacao: false,
	},
	{
		URI:                "/getClient/{id}",
		Metodo:             http.MethodGet,
		Funcao:             entity.GetClient,
		RequerAutenticacao: false,
	},
	{
		URI:                "/getAllClients",
		Metodo:             http.MethodGet,
		Funcao:             entity.GetAllClients,
		RequerAutenticacao: false,
	},
	{
		URI:                "/updateClient/{id}",
		Metodo:             http.MethodPut,
		Funcao:             entity.UpdateClient,
		RequerAutenticacao: false,
	},
}
