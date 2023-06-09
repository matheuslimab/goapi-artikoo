package rotas

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/entity"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             entity.Login,
	RequerAutenticacao: false,
}
