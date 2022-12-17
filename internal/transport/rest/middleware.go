package rest

import (
	"context"
	"net/http"
	"strings"

	"github.com/Levap123/trello-clone/pkg/jwt"
	"github.com/Levap123/trello-clone/pkg/webjson"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authString := r.Header.Get("Authorization")
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
