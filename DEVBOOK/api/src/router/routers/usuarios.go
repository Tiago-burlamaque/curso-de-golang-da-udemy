package routers

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuario = []Rota{
	{
		URI:              "/usuarios",
		Metodo:           http.MethodPost,
		Funcao:           controllers.CriarUsuario,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuarios",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuarios,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuario,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.AtualizandoUsuario,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodDelete,
		Funcao:           controllers.DeletarUsuario,
		RequerAutenticao: false,
	},
}
