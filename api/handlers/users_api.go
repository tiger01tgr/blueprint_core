package handlers

import (
	"backend/services/user"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func InitUsersRoutes(router chi.Router) {
	router.Route("/users", func(r chi.Router) {
		// Middlewares
		//r.Use(middleware.GoogleAuth)

		// Routes
		// r.Get("/all", GetUsers)
		r.Get("/", GetUserWithIdOrEmail)
		r.Get("/me", GetUserWithSelf)
		r.Post("/", CreateUser)
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUserWithIdOrEmail(w http.ResponseWriter, r *http.Request) {
	// If ID is given
	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err == nil {
		user, err := user.GetUserWithId(idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(user.String()))
		return
	}

	// If email is given
	email := r.FormValue("email")
	if email != "" {
		user, err := user.GetUserWithEmail(email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(user.String()))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
