package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/auth/sign-in", h.signIn)
	router.HandleFunc("/auth/sign-up", h.signUp)
	return router
}
