package server

import (
	"log"
	"net/http"
	"sispa-iam-api/internal/handler"
	"sispa-iam-api/internal/middleware"
	"sispa-iam-api/internal/service"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

const PORT = ":8090"

type Server struct {
	mux *http.ServeMux

	enforcerHandler *handler.EnforcerHandler
}

type NewServerInjectInput struct {
	CasdoorClient *casdoorsdk.Client
}

func NewServer(inject *NewServerInjectInput) *Server {
	casdoorService := service.NewCasdoorService(inject.CasdoorClient)

	enforcerService := service.NewEnforcerService(&service.ServiceInjectInput{
		CasdoorService: casdoorService,
	})
	enforcerHandler := handler.NewEnforcerHandler(enforcerService)

	s := &Server{
		mux:             http.NewServeMux(),
		enforcerHandler: enforcerHandler,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configura as rotas do servidor
func (s *Server) setupRoutes() {
	s.mux.Handle("/enforcer/enforce", middleware.AuthMiddleware(http.HandlerFunc(s.enforcerHandler.Enforce)))
}

// Start inicia o servidor
func (s *Server) Start() {
	log.Printf("Server starting on PORT %s", PORT)
	http.ListenAndServe(PORT, s.mux)
}
