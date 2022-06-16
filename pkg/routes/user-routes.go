package routes

import (
	"github.com/gorilla/mux"

	"github.com/blazingly-fast/axiomq-bookstore/pkg/controllers"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")

}
