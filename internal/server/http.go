package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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

func NewServer(cfg *configs.Config, h *httpHandler.UserHandler) (*Server, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "OK")
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetUser(w, r)
		case http.MethodPost:
			h.CreateUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
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
