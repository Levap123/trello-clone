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

type workspaceInput struct {
	Title string `json:"title,omitempty"`
	Logo  string `json:"logo,omitempty"`
}

// @Summary Create worskpace
// @Tags workspace
// @Description create workspace
// @ID create-workspace
// @Accept  json
// @Produce  json
// @Param input body workspaceInput true "credentials"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} postBody
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces [post]
func (h *Handler) createWorkspace(w http.ResponseWriter, r *http.Request) {
	var input workspaceInput
	webjson.ReadJSON(r, &input)
	userId := r.Context().Value("id").(int)
	workspaceId, err := h.service.Workspace.Create(input.Title, input.Logo, userId)
	if err != nil {
		h.logger.Err.Println(err.Error())
		webjson.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, postBody{workspaceId})
}

// @Summary Get worskpace by id
// @Tags workspace
// @Description get workspace
// @ID get-workspace
// @Accept  json
// @Produce  json
// @Param id path string true "Workspace ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} entity.Workspace
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{id} [get]
func (h *Handler) getWorkspacesById(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id").(int)
	workSpaceId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		webjson.JSONError(w, fmt.Errorf("Invalid URL"), http.StatusNotFound)
	}
	workSpace, err := h.service.Workspace.GetById(userId, workSpaceId)
	if err != nil {
		h.logger.Err.Println(err.Error())
		if errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, workSpace)
}

func (h *Handler) deleteWorkspaceById(w http.ResponseWriter, r *http.Request) {
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
