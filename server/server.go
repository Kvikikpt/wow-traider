package server

import (
	"github.com/Kvikikpt/wow-traider/db"
	"github.com/Kvikikpt/wow-traider/routes"
	"github.com/Kvikikpt/wow-traider/schedulers"
	"github.com/go-chi/chi"
	"net/http"
	"os"
)

type Server struct {
	config *Config
	router *chi.Mux
}

func New() *Server {
	return &Server{
		config: ReadConfig(),
		router: routes.NewRouter(),
	}
}

func (s *Server) Start() error {
	//Initialize schedulers
	schedulers.InitSchedulers()
	//starting db
	db.InitDb(s.config.databaseUrl)

	port := os.Getenv("PORT")
	println("Running Server on port:", port)
	return http.ListenAndServe(s.config.port, s.router)
}