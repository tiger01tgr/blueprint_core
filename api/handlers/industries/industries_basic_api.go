package handlers

import (
	industryService "backend/services/industries"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitIndustriesRoutes(router chi.Router) {
	router.Route("/industries", func(r chi.Router) {
		// Middlewares
		//r.Use(middleware.GoogleAuth)

		// Routes
		// r.Get("/all", GetEmployers)
		r.Get("/", GetIndustries)
		r.Post("/", CreateIndustries)
		r.Patch("/", EditIndustries)
		r.Delete("/", DeleteIndustries)
	})
}

func GetIndustries(w http.ResponseWriter, r *http.Request) {
	industries, err := industryService.GetAllIndustries()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Marshal the response to JSON
	jsonData, err := json.Marshal(industries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Write the response
	w.Write(jsonData)
}

func CreateIndustries(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name cannot be empty"))
		return
	}
	err := industryService.CreateIndustry(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func EditIndustries(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	if id == "" || name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id and name cannot be empty"))
		return
	}
	err := industryService.EditIndustry(id, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteIndustries(w http.ResponseWriter, r *http.Request) {
	// Get the request body
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id cannot be empty"))
		return
	}
	err := industryService.DeleteIndustry(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
