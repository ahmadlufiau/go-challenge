package controllers

import (
	"net/http"

	"github.com/ahmadlufiau/go-challenge/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome")

}
