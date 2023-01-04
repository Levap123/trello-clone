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
	webjson.SendJSON(w, map[string]int{"listId": listId})
}

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
