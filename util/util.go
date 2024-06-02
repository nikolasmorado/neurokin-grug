package util

import (
	"encoding/json"
	"net/http"
  t "neurokin/types"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type HandlerDependencies struct {
	Store t.Storage
}

func NewHandlerDependencies(store t.Storage) *HandlerDependencies {
	return &HandlerDependencies{
		Store: store,
	}
}

