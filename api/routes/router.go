package routes

import (
	employerHandler "backend/api/handlers/employers"
	industriesHandler "backend/api/handlers/industries"
	practiceHandler "backend/api/handlers/practice"
	usersHandler "backend/api/handlers/users"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var (
	router *chi.Router
	once   sync.Once
)

func GetRouter() *chi.Router {
	once.Do(InitRouter)
	return router
}

func InitRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	usersHandler.InitUsersRoutes(r)
	practiceHandler.InitPracticeRoutes(r)
	employerHandler.InitEmployersRoutes(r)
	industriesHandler.InitIndustriesRoutes(r)
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
