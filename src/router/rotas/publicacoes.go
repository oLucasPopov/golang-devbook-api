package rotas

import (
	"api/src/controllers"
	"net/http"
)

const (
	uriRotasPublicacoes      string = "/publicacoes"
	uriRotasPublicacoesComId string = "/publicacoes/{publicacaoId}"
)

var rotasPublicacoes = []Rota{
	{
		Uri:                uriRotasPublicacoes,
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                uriRotasPublicacoes,
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		Uri:                uriRotasPublicacoesComId,
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                uriRotasPublicacoesComId,
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                uriRotasPublicacoesComId,
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                uriRotasPublicacoesComId + "/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                uriRotasPublicacoesComId + "/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
