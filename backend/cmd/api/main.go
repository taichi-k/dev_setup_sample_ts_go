package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello from Go!")
	})

	readTimeout := 15
	writeTimeout := 15
	idleTimeout := 60

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}

	log.Println("Server running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
