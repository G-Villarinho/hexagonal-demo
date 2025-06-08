package http

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func ConfigureRoutes(userHandler *UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/users", userHandler.CreateUser)
		r.Post("/login", userHandler.Login)
	})

	return r
}
