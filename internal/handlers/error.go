package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(Response{
		Success: false,
		Error: map[string]string{
			"message": message,
		},
	})
}
