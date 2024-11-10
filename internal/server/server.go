package server

import (
	"log"
	"net/http"
	"sispa-iam-api/internal/handler"
)

const PORT = ":8090"

type Server struct {
	mux *http.ServeMux
}

// NewServer cria uma nova instância do servidor
func NewServer() *Server {
	s := &Server{mux: http.NewServeMux()}
	s.setupRoutes()
	return s
}

// setupRoutes configura as rotas do servidor
func (s *Server) setupRoutes() {
	s.mux.HandleFunc("/users", handler.GetUsers)
}

// Start inicia o servidor
func (s *Server) Start() {
	log.Printf("Server starting on PORT %s", PORT)
	http.ListenAndServe(PORT, s.mux)
}
