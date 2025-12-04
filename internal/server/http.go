package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrxacker/user_service/configs"
	httpHandler "github.com/mrxacker/user_service/internal/handlers/http"
)

const (
	ReadTimeout  = 5 * time.Second
	WriteTimeout = 10 * time.Second
	IdleTimeout  = 120 * time.Second
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *configs.Config, h *httpHandler.Handlers) (*Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "OK")
		if err != nil {
			return
		}
	})

	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/users", h.UserHandler.Routes())
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		IdleTimeout:  IdleTimeout,
	}

	return &Server{httpServer: srv}, nil
}

func (s *Server) Run() error {
	log.Printf("Server running at http://localhost%s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.httpServer.Shutdown(ctx)
}
