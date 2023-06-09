package router

import (
	"github.com/matheuslimab/artikoo/api/router/rotas"

	"github.com/gorilla/mux"
)

// gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
