package middleware

import (
	"context"
	"net/http"
	"sispa-iam-api/internal/service"
	"strings"
)

type contextKey string

const (
	AuthorizationHeader = "Authorization"
	TokenContextKey     = contextKey("token")
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(AuthorizationHeader)

		if authHeader == "" {
			service.StdHttpError(w, &service.ErrorHandlerInput{
				Message:        []string{"Authorization header missing"},
				HttpStatusCode: http.StatusBadRequest,
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			service.StdHttpError(w, &service.ErrorHandlerInput{
				Message:        []string{"Invalid authorization header format"},
				HttpStatusCode: http.StatusBadRequest,
			})
			return
		}

		token := parts[1]
		ctx := context.WithValue(r.Context(), TokenContextKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(TokenContextKey).(string)
	return token, ok
}
