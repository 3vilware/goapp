package routes

import (
	"github.com/gorilla/mux"
	"github.com/goapp/controlers"
)

func SetLoginRouter(router *mux.Router){
	router.HandleFunc("/api/login", controlers.Login).Methods("POST")
}