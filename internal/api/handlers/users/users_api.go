package auth

import (
	"backend/internal/services/user"
	"backend/internal/api/middleware"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func InitUsersRoutes(router chi.Router) {
	router.Route("/users", func(r chi.Router) {
		// Middlewares
		r.Use(middleware.GoogleAuth)

		// Routes
		// r.Get("/", GetUsers)
		r.Get("/{id}", GetUserWithId)
		r.Get("/{email}", GetUserWithEmail)
		r.Get("/me", GetUserWithSelf)
		r.Post("/", CreateUser)
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUserWithId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := user.GetUserWithId(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(user.String()))
}

func GetUserWithEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	user, err := user.GetUserWithEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(user.String()))
}

func GetUserWithSelf(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int)
	user, err := user.GetUserWithId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(user.String()))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	middlename := r.FormValue("middlename")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	typeOfUser := r.FormValue("typeOfUser")
	err := user.CreateUser(firstname, middlename, lastname, email, typeOfUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
