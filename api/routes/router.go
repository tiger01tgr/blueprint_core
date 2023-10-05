package routes

import (
	"backend/api/handlers"
	"fmt"
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
	handlers.InitUsersRoutes(r)

	http.ListenAndServe(":3000", r)
	fmt.Println("Server is running on port 3000")
}
