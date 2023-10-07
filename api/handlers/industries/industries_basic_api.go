package handlers

import (
	"backend/db/models"
	industryService "backend/services/employers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitEmployersRoutes(router chi.Router) {
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
	// Get the request body
	var industry models.Industry
	err := json.NewDecoder(r.Body).Decode(&industry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Create the industry
	err = industryService.CreateIndustry(industry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EditIndustries(w http.ResponseWriter, r *http.Request) {
	// Get the request body
	var industry models.Industry
	err := json.NewDecoder(r.Body).Decode(&industry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Edit the industry
	err = industryService.EditIndustry(industry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteIndustries(w http.ResponseWriter, r *http.Request) {
	// Get the request body
	var industry models.Industry
	err := json.NewDecoder(r.Body).Decode(&industry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Delete the industry
	err = industryService.DeleteIndustry(industry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
