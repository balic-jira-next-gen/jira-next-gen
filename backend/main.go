package main

import (
	"log"
	"net/http"
	"os"

	"github.com/concept/jira_next_gen/backend/internal/api"
	"github.com/concept/jira_next_gen/backend/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize Store
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://jira_user:jira_password@localhost:5432/jira_next_gen?sslmode=disable"
	}

	s, err := store.NewPostgresStore(dbURL)
	if err != nil {
		log.Fatalf("Error initializing store: %v", err)
	}

	// Initialize API Handler
	h := api.NewHandler(s)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CorsHandler(middleware.CorsOptions{
		AllowedOrigins: []string{"*"}, // Allow all for dev
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Jira Next Gen Backend API is running"))
	})

	// Mount API routes
	r.Route("/api", func(r chi.Router) {
		h.RegisterRoutes(r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	http.ListenAndServe(":"+port, r)
}
