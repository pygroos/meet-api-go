package main

import (
	"meet-api/routes"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8081",
		Handler:      route.InitRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
