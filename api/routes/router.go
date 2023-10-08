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
	usersHandler.InitUsersRoutes(r)
	practiceHandler.InitPracticeRoutes(r)
	employerHandler.InitEmployersRoutes(r)
	industriesHandler.InitIndustriesRoutes(r)

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
