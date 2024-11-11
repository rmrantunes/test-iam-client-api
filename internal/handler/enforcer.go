package handler

import (
	"encoding/json"
	"net/http"
	"sispa-iam-api/internal/middleware"
	"sispa-iam-api/internal/service"
)

type EnforcerHandler struct {
	enforcerService *service.EnforcerService
}

func NewEnforcerHandler(enforcerService *service.EnforcerService) *EnforcerHandler {
	return &EnforcerHandler{
		enforcerService: enforcerService,
	}
}

type EnforcerHandlerData struct {
	resource             string
	action               string
	relationObjectUserId string
}

func (s *EnforcerHandler) Enforce(w http.ResponseWriter, r *http.Request) {
	token, ok := middleware.GetTokenFromContext(r.Context())

	if !ok {
		http.Error(w, "Token not found in context", http.StatusUnauthorized)
	}

	var data EnforcerHandlerData
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "error while decode body JSON", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()

	permissionId := query.Get("permissionId")

	if permissionId == "" {
		service.StdHttpError(w, &service.ErrorHandlerInput{
			Message:        []string{"permissionId missing"},
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := s.enforcerService.Enforce(&service.EnforceInput{
		AccessToken:          token,
		PermissionId:         permissionId,
		Resource:             data.resource,
		Action:               data.action,
		RelationObjectUserId: data.relationObjectUserId,
	})

	if err != nil {
		service.StdHttpError(w, &service.ErrorHandlerInput{
			Message:        []string{err.Error()},
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
