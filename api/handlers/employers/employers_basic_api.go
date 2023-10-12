package handlers

import (
	// "backend/api/middleware"
	"backend/db/models"
	employersService "backend/services/employers"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitEmployersRoutes(router chi.Router) {
	router.Route("/api/employers", func(r chi.Router) {
		// Public routes
		r.Group(func(r chi.Router) {
			r.Get("/", GetEmployers)
		})

		// Admin routes
		r.Group(func(r chi.Router) {
			// middleware.SuperAdminAuth(r)
			r.Post("/", CreateEmployer)
			r.Patch("/", EditEmployer)
			r.Delete("/", DeleteEmployer)
		})
	})
}

type EmployerResponse struct {
	ID         uint64
	Name       string
	Logo       string
	Industry   string
	IndustryId uint64
	Deleted   bool
}

func GetEmployers(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("all") == "true" {
		// Return all employers
		employers, err := employersService.GetAllEmployers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		response := employersToResponseHelper(employers)

		// Marshal the response to JSON
		jsonData, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Set the content type header and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

	// var req GetEmployersRequest

	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&req); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Invalid request payload"))
	// 	return
	// }

}

func CreateEmployer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	logo, _, err := r.FormFile("logo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer logo.Close()

	industryId := r.FormValue("industryId")

	if name == "" || industryId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	err = employersService.CreateEmployer(name, industryId, logo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Converts []models.Employer to []EmployerResponse
func employersToResponseHelper(employers []models.Employer) []EmployerResponse {
	var responses []EmployerResponse
	for _, employer := range employers {
		response := EmployerResponse{
			ID:         employer.ID,
			Name:       employer.Name,
			Logo:       employer.Logo,
			Industry:   employer.Industry,
			IndustryId: employer.IndustryId,
			Deleted:   employer.Deleted,
		}
		responses = append(responses, response)
	}
	return responses
}

func EditEmployer(w http.ResponseWriter, r *http.Request) {
	var logo multipart.File
	if r.FormValue("isLogoUpdate") == "true" {
		err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to parse form or form is too large"))
			return
		}
		logo, _, err = r.FormFile("logo")
		defer logo.Close()
		if err != nil && err != http.ErrMissingFile {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}

	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	name := r.FormValue("name")

	industryId := r.FormValue("industryId")

	err := employersService.UpdateEmployer(id, name, industryId, logo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteEmployer(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	err := employersService.DeleteEmployer(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
