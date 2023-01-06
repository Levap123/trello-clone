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

type boardInput struct {
	Title      string `json:"title,omitempty"`
	Background string `json:"background,omitempty"`
}

// @Summary Create board
// @Tags board
// @Description create board
// @ID create-board
// @Accept  json
// @Produce  json
// @Param input body boardInput true "credentials"
// @Param workspaceId path string true "Workspace ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} postBody
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards [post]
func (h *Handler) createBoard(w http.ResponseWriter, r *http.Request) {
	defaultLists := []string{"To Do", "Doing", "Finished"}
	var input boardInput
	webjson.ReadJSON(r, &input)
	workspaceId, err := strconv.Atoi(mux.Vars(r)["workspaceId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(workspaceId)
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
	for i := 0; i < len(defaultLists); i++ {
		_, err := h.service.List.Create(defaultLists[i], userId, workspaceId, boardId)
		if err != nil {
			h.logger.Err.Println(err)
			webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	webjson.SendJSON(w, postBody{boardId})
}

// @Summary Get board by id
// @Tags board
// @Description get board
// @ID get-board
// @Accept  json
// @Produce  json
// @Param workspaceId path string true "Workspace ID"
// @Param id path string true "Board ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} entity.Board
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards/{id} [get]
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
		if errors.Is(err, errs.ErrInvalidBoard) || errors.Is(err, errs.ErrInvalidWorkspace) || errors.Is(err, errs.ErrForeignKeyFailed) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, board)
}

// @Summary Get boards by workspace id
// @Tags board
// @Description get boards
// @ID get-boards
// @Accept  json
// @Produce  json
// @Param workspaceId path string true "Workspace ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {array} entity.Board
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards [get]
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
		if errors.Is(err, errs.ErrInvalidWorkspace) || errors.Is(err, errs.ErrForeignKeyFailed) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusBadRequest)
		return
	}
	webjson.SendJSON(w, boards)
}

// @Summary Delete board by id
// @Tags board
// @Description delete board
// @ID delete-board
// @Accept  json
// @Produce  json
// @Param workspaceId path string true "Workspace ID"
// @Param id path string true "Board ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} postBody
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards/{id} [delete]
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
		if errors.Is(err, errs.ErrInvalidBoard) || errors.Is(err, errs.ErrInvalidWorkspace) || errors.Is(err, errs.ErrForeignKeyFailed) {
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]int{"boardId": boardId})
}
