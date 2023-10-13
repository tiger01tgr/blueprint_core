package handlers

import (
	"backend/api/middleware"
	"backend/services/user"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func InitUsersRoutes(router chi.Router) {
	router.Route("/api/users", func(r chi.Router) {
		// Middlewares
		//r.Use(middleware.GoogleAuth)

		// Admin Routes
		r.Get("/", GetUser)

		// User Routes
		r.Group(func(r chi.Router) {
			middleware.UserAuth(r)
			r.Get("/me", GetUserWithSelf)
			r.Post("/", CreateUser)
		})
	})
}

type UserResponse struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	UserType   string `json:"type_of_user"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// If ID is given
	id := r.FormValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		user, err := user.GetUserWithId(idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
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
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(user.String()))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func GetUserWithSelf(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)
	user, err := user.GetUserWithEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response := UserResponse{
		ID:         int64(user.ID),
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		Email:      user.Email,
		UserType:   user.UserType,
	}
	w.Header().Set("Content-Type", "application/json")

	// Encode the response as JSON and write it to the response writer
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	return
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
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
