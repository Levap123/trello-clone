package rest

import (
	"errors"
	"net/http"
	"strconv"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/Levap123/trello-clone/pkg/webjson"
	"github.com/gorilla/mux"
)

type boardInput struct {
	Title      string `json:"title,omitempty"`
	Background string `json:"background,omitempty"`
}

func (h *Handler) createBoard(w http.ResponseWriter, r *http.Request) {
	var input boardInput
	webjson.ReadJSON(r, &input)
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	boardId, err := h.service.Board.Create(input.Title, input.Background, userId, workspaceId)
	if err != nil {
		h.logger.Err.Println(err)
		if errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]int{"boardId": boardId})
}

func (h *Handler) getBoardById(w http.ResponseWriter, r *http.Request) {
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	boardId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	board, err := h.service.Board.GetById(userId, boardId, workspaceId)
	if err != nil {
		h.logger.Err.Println(err)
		if errors.Is(err, errs.ErrInvalidBoard) || errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, board)
}

func (h *Handler) getBoardByWorkspaceId(w http.ResponseWriter, r *http.Request) {
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	boards, err := h.service.Board.GetByWorkspaceId(userId, workspaceId)
	if err != nil {
		h.logger.Err.Println(err)
		if errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusBadRequest)
		return
	}
	webjson.SendJSON(w, boards)
}

func (h *Handler) deleteBoardById(w http.ResponseWriter, r *http.Request) {
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	boardId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	boardId, err = h.service.Board.DeleteById(userId, workspaceId, boardId)
	if err != nil {
		h.logger.Err.Println(err)
		if errors.Is(err, errs.ErrInvalidBoard) || errors.Is(err, errs.ErrInvalidWorkspace) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]int{"boardId": boardId})
}
