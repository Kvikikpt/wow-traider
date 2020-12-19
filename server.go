package main

import (
	"context"
	"github.com/Kvikikpt/wow-traider/db"
	"github.com/Kvikikpt/wow-traider/schedulers"
	"github.com/carlescere/scheduler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//loading .env files
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	//starting db
	var dbErr error
	db.Pool, dbErr = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if dbErr != nil {
		log.Fatalf("Unable to connect to database: %v", dbErr)
	}
	defer db.Pool.Close(context.Background())

	//starting cron jobs
	_, _ = scheduler.Every(2).Seconds().NotImmediately().Run(schedulers.TestJob)

	//set routes
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", index)

	//start server
	port := os.Getenv("PORT")
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