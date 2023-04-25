package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaTransactions = []Rota{
	{
		URI:                "/createTransaction",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateTransaction,
		RequerAutenticacao: false,
	},
}
