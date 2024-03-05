package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BoruTamena/RssAggergetor/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// hold connection to the database
type apiConfig struct {
	db *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port is not found in the enviroment")
	}

	// importing database connection
	db_url := os.Getenv("DBURL")
	if db_url == "" {
		log.Fatal("DBURL is not fond in the enviroment ")
	}

	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("can't connect to the database ", err)
	}

	apicfg := apiConfig{
		db: database.New(conn),
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
	v1Router.Post("/user", apicfg.handlerCreateUser)
	v1Router.Get("/user", apicfg.handleGetUser)

	router.Mount("/v1", v1Router)

	// connecting router with http server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("server starting on port %v", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
