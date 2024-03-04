package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("port is not found n the eviroment")
	}

	fmt.Println("port--", port)

	// creating a router
	router := chi.NewRouter()
	// configuring the router
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Links"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// mapping http handler with specific route path and methods

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)

	router.Mount("/v1", v1Router)

	// connecting router with http server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("server starting on port %v", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
