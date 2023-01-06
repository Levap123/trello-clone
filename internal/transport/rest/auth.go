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

type tokenOutput struct {
	Token string `json:"token,omitempty"`
}

type signUpBody struct {
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

// @Summary sign up
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body signUpBody true "credentials"
// @Success 200 {object} postBody
// @Failure default {object} webjson.errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var body signUpBody
	webjson.ReadJSON(r, &body)
	id, err := h.service.Auth.CreateUser(body.Email, body.Name, body.Password)
	if err != nil {
		if errors.Is(err, errs.ErrUniqueUser) {
			h.logger.Err.Printf("%v\n", err)
			webjson.JSONError(w, err, http.StatusBadRequest)
			return
		}
		h.logger.Err.Printf("%v\n", err)
		webjson.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, postBody{id})
}

// @Summary sign in
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInBody true "credentials"
// @Success 200 {object} tokenOutput
// @Failure default {object} webjson.errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var body signInBody
	webjson.ReadJSON(r, &body)
	token, err := h.service.Auth.GetUser(body.Email, body.Password)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, errs.ErrInvalidEmail) || errors.Is(err, errs.ErrPasswordIncorrect) {
			status = http.StatusBadRequest
		}
		webjson.JSONError(w, err, status)
		return
	}
	webjson.SendJSON(w, tokenOutput{token})
}
