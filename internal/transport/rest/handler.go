package rest

import (
	"net/http"

	"github.com/Levap123/trello-clone/internal/service"
	"github.com/Levap123/trello-clone/pkg/logger"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
	logger  *logger.Logger
}

func NewHandler(service *service.Service, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	{
		router.HandleFunc("/auth/sign-in", h.signIn).Methods("POST")
		router.HandleFunc("/auth/sign-up", h.signUp).Methods("POST")
	}
	api := router.PathPrefix("/api").Subrouter()
	api.Use(h.userIdentity)
	workspace := api.PathPrefix("/workspaces").Subrouter()
	{
		workspace.HandleFunc("", h.createWorkspace).Methods("POST")
		workspace.HandleFunc("/{id}", h.getWorkspacesById).Methods("GET")
		workspace.HandleFunc("/{id}", h.deleteWorkspaceById).Methods("DELETE")
		boards := workspace.PathPrefix("/{workspaceId}/boards").Subrouter()
		{
			boards.HandleFunc("", h.createBoard).Methods("POST")
			boards.HandleFunc("", h.getBoardByWorkspaceId).Methods("GET")
			boards.HandleFunc("/{id}", h.getBoardById).Methods("GET")
			boards.HandleFunc("/{id}", h.deleteBoardById).Methods("DELETE")
			lists := boards.PathPrefix("/{boardId}/lists").Subrouter()
			{
				lists.HandleFunc("", h.createList).Methods("POST")
			}
		}
	}

	return router
}
