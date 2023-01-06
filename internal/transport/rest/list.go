package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/Levap123/trello-clone/pkg/webjson"
	"github.com/gorilla/mux"
)

type listInput struct {
	Title string `json:"title,omitempty"`
}

// @Summary Create list
// @Tags list
// @Description create list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body listInput true "credentials"
// @Param workspaceId path string true "Workspace ID"
// @Param boardId path string true "Board ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} postBody
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards/{boardId}/lists [post]
func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	boardId, err := strconv.Atoi(mux.Vars(r)["boardId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	var input listInput
	webjson.ReadJSON(r, &input)
	listId, err := h.service.List.Create(input.Title, userId, workspaceId, boardId)
	if err != nil {
		h.logger.Err.Println(err)
		if errors.Is(err, errs.ErrInvalidBoard) || errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, postBody{listId})
}

// @Summary Get lists by board
// @Tags list
// @Description get lists
// @ID get-lists
// @Accept  json
// @Produce  json
// @Param workspaceId path string true "Workspace ID"
// @Param boardId path string true "Board ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {array} entity.List
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards/{boardId}/lists [get]
func (h *Handler) getByBoardId(w http.ResponseWriter, r *http.Request) {
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	boardId, err := strconv.Atoi(mux.Vars(r)["boardId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	lists, err := h.service.List.GetByBoardId(userId, workspaceId, boardId)
	if err != nil {
		h.logger.Err.Println(err)
		if errors.Is(err, errs.ErrInvalidBoard) || errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "no rows in result set") {
			webjson.JSONError(w, fmt.Errorf("list with this id does not exist"), http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, lists)
}

// @Summary Get list by list id
// @Tags list
// @Description get list
// @ID get-list
// @Accept  json
// @Produce  json
// @Param workspaceId path string true "Workspace ID"
// @Param boardId path string true "Board ID"
// @Param id path string true "Board ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} entity.List
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards/{boardId}/lists/{id} [get]
func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	boardId, err := strconv.Atoi(mux.Vars(r)["boardId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	listId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	list, err := h.service.List.GetById(userId, workspaceId, boardId, listId)
	if err != nil {
		h.logger.Err.Println(err)
		if errors.Is(err, errs.ErrForeignKeyFailed) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "no rows in result set") {
			webjson.JSONError(w, fmt.Errorf("list with this id does not exist"), http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, list)
}

// @Summary delete list by list id
// @Tags list
// @Description delete list
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param workspaceId path string true "Workspace ID"
// @Param boardId path string true "Board ID"
// @Param id path string true "Board ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} postBody
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards/{boardId}/lists/{id} [delete]
func (h *Handler) deleteListById(w http.ResponseWriter, r *http.Request) {
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	boardId, err := strconv.Atoi(mux.Vars(r)["boardId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	listId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	listId, err = h.service.List.DeleteById(userId, workspaceId, boardId, listId)
	if err != nil {
		h.logger.Err.Println(err)
		if errors.Is(err, errs.ErrForeignKeyFailed) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "no rows in result set") {
			webjson.JSONError(w, fmt.Errorf("list with this id does not exist"), http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]int{"listId": listId})
}
