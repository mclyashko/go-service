package http

import (
	"context"
	"net/http"

	"github.com/mclyashko/go-service/internal/http/handler/hello_world"
	"github.com/mclyashko/go-service/internal/http/handler/joke"
)

type Server struct {
	srv *http.Server
}

func NewServer(
	addr string,
	hw hello_world.Handler,
	j joke.Handler,
) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", hw.Hello)
	mux.HandleFunc("GET /joke", j.Joke)

	return &Server{
		srv: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
