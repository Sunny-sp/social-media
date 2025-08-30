package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func ResponseError(w http.ResponseWriter, statusCode int, message ...string) {
	msg := http.StatusText(statusCode)
	if len(message) > 0 {
		msg = message[0] // custom override
	}
	ResponseJSON(w, statusCode, map[string]string{"error": msg})
}
