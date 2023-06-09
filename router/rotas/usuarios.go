package rotas

import (
	"net/http"

	"github.com/matheuslimab/artikoo/api/internal/entity"
	"github.com/matheuslimab/artikoo/api/src/helpers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             entity.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             entity.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodGet,
		Funcao:             entity.BuscaUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodPut,
		Funcao:             entity.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             entity.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:             helpers.GET,
		Funcao:             entity.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
