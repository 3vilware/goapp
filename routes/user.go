package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/goapp/controlers"
)

// Ruta para registro de usuario
func SetUserRouter(router *mux.Router){
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controlers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}