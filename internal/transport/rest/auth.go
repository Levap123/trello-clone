package rest

import (
	"errors"
	"net/http"

	errs "github.com/Levap123/trello-clone/pkg/errors"
	"github.com/Levap123/trello-clone/pkg/webjson"
)

type signInBody struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var body signInBody
	webjson.ReadJSON(r, &body)
	id,err=
}

type signUpBody struct {
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var body signUpBody
	webjson.ReadJSON(r, &body)
	id, err := h.service.Auth.CreateUser(body.Email, body.Name, body.Password)
	if err != nil {
		if errors.Is(err, errs.UniqueErr) {
			h.logger.Err.Printf("%v\n", err)
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		h.logger.Err.Printf("%v\n", err)
		webjson.JSONError(w, errs.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]int{"id": id})
}
