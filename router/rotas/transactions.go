package rotas

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/entity"
)

var rotaTransactions = []Rota{
	{
		URI:                "/createTransaction",
		Metodo:             http.MethodPost,
		Funcao:             entity.CreateTransaction,
		RequerAutenticacao: false,
	},
}
