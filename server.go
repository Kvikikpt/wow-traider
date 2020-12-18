package main

import (
	"github.com/Kvikikpt/wow-traider/schedulers"
	"github.com/carlescere/scheduler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	//loading .env files
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	//starting cron jobs
	scheduler.Every(2).Seconds().NotImmediately().Run(schedulers.TestJob)

	//set routs
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", index)

	//start server
	port := ":8080"
	println("Running server on port:", port)
	serverErr := http.ListenAndServe(port, router)
	if serverErr != nil {
		log.Fatal("ListenAndServe", serverErr)
	}
}

//index page
func index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Index"))
}