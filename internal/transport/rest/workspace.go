package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/Levap123/trello-clone/pkg/webjson"
	"github.com/gorilla/mux"
)

type WorksSpaceInput struct {
	Title string `json:"title,omitempty"`
	Logo  string `json:"logo,omitempty"`
}

func (h *Handler) CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	var input WorksSpaceInput
	webjson.ReadJSON(r, &input)
	userId := r.Context().Value("id").(int)
	id, err := h.service.Workspace.Create(input.Title, input.Logo, userId)
	if err != nil {
		h.logger.Err.Println(err.Error())
		webjson.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]int{"workspaceId": id})
}

func (h *Handler) GetWorkspacesById(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id").(int)
	workSpaceId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		webjson.JSONError(w, fmt.Errorf("Invalid URL"), http.StatusNotFound)
	}
	worksSpace, err := h.service.Workspace.GetById(userId, workSpaceId)
	if err != nil {
		h.logger.Err.Println(err.Error())
		if errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, worksSpace)
}

func (h *Handler) DeleteWorkspaceById(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id").(int)
	fmt.Println(userId)
	workSpaceId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		webjson.JSONError(w, fmt.Errorf("Invalid URL"), http.StatusNotFound)
	}
	workSpaceId, err = h.service.Workspace.DeleteById(userId, workSpaceId)
	if err != nil {
		h.logger.Err.Println(err.Error())
		if errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]int{"workspaceId": workSpaceId})
}
