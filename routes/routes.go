package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux{
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", Index)

	return router
}
