package handler

import (
	"encoding/json"
	"net/http"
	service_users "sispa-iam-api/internal/service"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := service_users.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
