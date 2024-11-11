package middleware

import (
	"net/http"
	"sispa-iam-api/internal/service"
)

func Method(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if method == "" {
			service.StdHttpError(w, &service.ErrorHandlerInput{
				Message:        []string{"Method middleware value not set"},
				HttpStatusCode: http.StatusInternalServerError,
			})
			return
		}

		if r.Method != method {
			service.StdHttpError(w, &service.ErrorHandlerInput{
				Message:        []string{"Method not allowed"},
				HttpStatusCode: http.StatusMethodNotAllowed,
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
