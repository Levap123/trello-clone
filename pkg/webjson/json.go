package webjson

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func SendJSON(w http.ResponseWriter, input any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(input)
}

