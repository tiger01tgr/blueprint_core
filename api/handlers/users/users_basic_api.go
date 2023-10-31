package handlers

import (
	"backend/api/middleware"
	"backend/services/user"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitUsersRoutes(router chi.Router) {
	router.Route("/api/users", func(r chi.Router) {
		// Middlewares
		//r.Use(middleware.GoogleAuth)

		// Admin Routes
		// r.Get("/", GetUser)
		r.Post("/", CreateUser)

		// User Routes
		r.Group(func(r chi.Router) {
			middleware.UserAuth(r)
			r.Get("/me", GetUserWithSelf)
			r.Get("/me/profile", GetUserProfileWithSelf)
			r.Patch("/me/profile", PatchUserProfile)
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

type UserProfileResponse struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	UserType   string `json:"type_of_user"`
	School     string `json:"school"`
	Major      string `json:"major"`
	Employer   string `json:"employer"`
	Position   string `json:"position"`
	Phone      string `json:"phone"`
	Resume     string `json:"resume"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	// If ID is given
// 	id := r.FormValue("id")
// 	idInt, err := strconv.ParseInt(id, 10, 64)
// 	if err == nil {
// 		user, err := user.GetUserWithId(idInt)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		w.Write([]byte(user.String()))
// 		return
// 	}

// 	// If email is given
// 	email := r.FormValue("email")
// 	if email != "" {
// 		id, err := user.GetUserIdWithEmail(email)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		w.Write([]byte(user.String()))
// 		return
// 	}

// 	w.WriteHeader(http.StatusBadRequest)
// }

func GetUserWithSelf(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int64)
	user, err := user.GetUserWithId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response := UserResponse{
		ID:         user.ID,
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

func GetUserProfileWithSelf(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int64)
	user, err := user.GetUserProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response := UserProfileResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		Email:      user.Email,
		UserType:   user.UserType,
		School:     user.School,
		Major:      user.Major,
		Employer:   user.Employer,
		Position:   user.Position,
		Phone:      user.Phone,
		Resume:     user.Resume,
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

func PatchUserProfile(w http.ResponseWriter, r *http.Request) {
	var resume *multipart.File
	if r.FormValue("isResumeUpdate") == "true" {
		err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to parse form or form is too large"))
			return
		}
		resumeFile, _, err := r.FormFile("resume")
		defer resumeFile.Close()
		if err != nil && err != http.ErrMissingFile {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if err == http.ErrMissingFile {
			resume = nil
		} else {
			resume = &resumeFile
		}
	}

	id := r.Context().Value("id").(int64)
	firstname := r.FormValue("firstname")
	middlename := r.FormValue("middlename")
	lastname := r.FormValue("lastname")
	school := r.FormValue("school")
	major := r.FormValue("major")
	employer := r.FormValue("employer")
	position := r.FormValue("position")
	phone := r.FormValue("phone")
	err := user.UpdateUserProfile(id, firstname, middlename, lastname, school, major, employer, position, phone, resume)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
