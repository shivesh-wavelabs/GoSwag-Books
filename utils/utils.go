package utils

import (
	"go_backend/models"
	"encoding/json"
	"net/http"
)

// Utility to send error
func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

// Utility to send success
func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
