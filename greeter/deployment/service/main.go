package main

import (
	"log"
	"net/http"
	"os"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	srv := handlers.NewServer()
	handler := gen.Handler(gen.NewStrictHandler(srv, nil))

	log.Printf("greeter listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
