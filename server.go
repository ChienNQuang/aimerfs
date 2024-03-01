package main

import (
	"net/http"

	"github.com/ChienNQuang/aimerfs/fs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type server struct {
	t *fs.Tree
}

func newServer() *server {
	return &server{
		t: fs.NewTree(),
	}
}

func (s *server) Serve() error {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	s.Routes(r)

	return http.ListenAndServe(":8080", r)
}
