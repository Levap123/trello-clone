package rest

import (
	"net/http"
	"strconv"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/Levap123/trello-clone/pkg/webjson"
	"github.com/gorilla/mux"
)

type cardInput struct {
	Title string `json:"title,omitempty"`
}

func (h *Handler) createCard(w http.ResponseWriter, r *http.Request) {
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
	listId, err := strconv.Atoi(mux.Vars(r)["listId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	var input cardInput
	webjson.ReadJSON(r, &input)
	cardId, err := h.service.Card.Create(input.Title, userId, workspaceId, boardId, listId)
	if err != nil {
		h.logger.Err.Println(err)
		webjson.JSONError(w, err, http.StatusBadRequest)
		return
	}
	webjson.SendJSON(w, postBody{cardId})
}

func (h *Handler) getCardsByListId(w http.ResponseWriter, r *http.Request) {
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
	listId, err := strconv.Atoi(mux.Vars(r)["listId"])
	if err != nil {
		webjson.JSONError(w, errs.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userId := r.Context().Value("id").(int)
	cards, err := h.service.Card.GetByListId(userId, workspaceId, boardId, listId)
	if err != nil {
		h.logger.Err.Println(err)
		webjson.JSONError(w, err, http.StatusBadRequest)
		return
	}
	webjson.SendJSON(w, cards)
}
