package main

import (
	"net/http"
	"time"
	"meet-api/routes"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      route.InitRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}