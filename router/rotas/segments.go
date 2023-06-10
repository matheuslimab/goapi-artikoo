package rotas

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/entity"
)

var rotaSegments = []Rota{
	{
		URI:                "/getAllSegments",
		Metodo:             http.MethodGet,
		Funcao:             entity.GetAllSegments,
		RequerAutenticacao: false,
	},
}
