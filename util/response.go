package util

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func ServerError(w http.ResponseWriter, v any) error {
	return WriteJSON(w, http.StatusInternalServerError, v)
}
