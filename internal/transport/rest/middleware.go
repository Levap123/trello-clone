package rest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Levap123/trello-clone/pkg/jwt"
	"github.com/Levap123/trello-clone/pkg/webjson"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authString := r.Header.Get("Authorization")
		if len(strings.Split(authString, "Bearer ")) < 2 {
			webjson.JSONError(w, fmt.Errorf("auth header not found"), http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(authString, "Bearer ")[1]
		id, err := jwt.ParseToken(tokenString)
		if err != nil {
			webjson.JSONError(w, err, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
