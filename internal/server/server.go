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

	userHandler *handler.UserHandler
}

// NewServer cria uma nova instância do servidor
func NewServer(casdoorClient *casdoorsdk.Client) *Server {
	casdoorService := service.NewCasdoorService(casdoorClient)

	userService := service.NewUserService(casdoorService)
	userHandler := handler.NewUserHandler(userService)

	s := &Server{mux: http.NewServeMux(), userHandler: userHandler}
	s.setupRoutes()
	return s
}

// setupRoutes configura as rotas do servidor
func (s *Server) setupRoutes() {
	s.mux.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(s.userHandler.GetUsers)))
}

// Start inicia o servidor
func (s *Server) Start() {
	log.Printf("Server starting on PORT %s", PORT)
	http.ListenAndServe(PORT, s.mux)
}
