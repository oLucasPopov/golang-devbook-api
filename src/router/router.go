package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar creates and returns a new mux.Router instance.
//
// This function does not take any parameters.
// It returns a pointer to a mux.Router.
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
