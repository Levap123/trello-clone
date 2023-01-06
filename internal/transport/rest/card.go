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

// @Summary Create card
// @Tags card
// @Description create card
// @ID create-card
// @Accept  json
// @Produce  json
// @Param input body cardInput true "credentials"
// @Param workspaceId path string true "Workspace ID"
// @Param boardId path string true "Board ID"
// @Param listId path string true "List ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} postBody
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards/{boardId}/lists/{listId}/cards [post]
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

// @Summary Get cards by list id
// @Tags card
// @Description get cards
// @ID get-cards
// @Accept  json
// @Produce  json
// @Param workspaceId path string true "Workspace ID"
// @Param boardId path string true "Board ID"
// @Param listId path string true "list ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {array} entity.Card
// @Failure default {object} webjson.errorResponse
// @Router /api/workspaces/{workspaceId}/boards/{boardId}/lists/{listId}/cards [get]
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
