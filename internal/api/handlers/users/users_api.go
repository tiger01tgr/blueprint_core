package auth

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func InitUsersRoutes(router chi.Router) {
	router.Route("/users", func(r chi.Router) {
		r.Get("/", GetUsers)
		r.Get("/{id}", GetUserWithId)
		r.Get("/{email}", GetUserWithEmail)
		r.Post("/", CreateUser)
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	
}

func GetUserWithId(w http.ResponseWriter, r *http.Request) {

}

func GetUserWithEmail(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}