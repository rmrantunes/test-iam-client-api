package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sispa-iam-api/internal/middleware"
	service_users "sispa-iam-api/internal/service"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	token, ok := middleware.GetTokenFromContext(r.Context())

	if !ok {
		http.Error(w, "Token not found in context", http.StatusUnauthorized)
	}

	fmt.Printf("Token received %s\n", token)

	users := service_users.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
