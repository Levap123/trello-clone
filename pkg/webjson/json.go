package webjson

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message,omitempty"`
}

func JSONError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorResponse{err.Error()})
}

func SendJSON(w http.ResponseWriter, input any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(input)
}

func ReadJSON(r *http.Request, output any) {
	json.NewDecoder(r.Body).Decode(output)
	defer r.Body.Close()
}
