package controllers

import (
	"net/http"

	"github.com/blazingly-fast/axiomq-bookstore/pkg/models"
)

var User models.User

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}
