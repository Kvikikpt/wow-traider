package main

import (
	"github.com/Kvikikpt/wow-traider/server"
	"log"
)

func main() {
	//start server
	s := server.New()
	if serverErr := s.Start(); serverErr != nil {
		log.Fatal("ListenAndServe", serverErr)
	}
}