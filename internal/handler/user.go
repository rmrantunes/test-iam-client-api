package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sispa-iam-api/internal/middleware"
	"sispa-iam-api/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (s *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	token, ok := middleware.GetTokenFromContext(r.Context())

	if !ok {
		service.StdHttpError(w, &service.ErrorHandlerInput{
			Message:        []string{"Token not found in header"},
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	fmt.Printf("Token received %s\n", token)

	users := s.userService.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
