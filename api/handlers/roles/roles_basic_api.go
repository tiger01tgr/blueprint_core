package handlers

import (
	rolesService "backend/services/roles"
	"encoding/json"

	//"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRolesRoute(router chi.Router) {
	router.Route("/api/roles", func(r chi.Router) {
		// Middlewares
		//r.Use(middleware.GoogleAuth)

		// Routes
		// r.Get("/all", GetEmployers)
		r.Get("/", GetIndustries)
		r.Post("/", CreateRole)
		r.Patch("/", EditRole)
		r.Delete("/", DeleteRole)
	})
}

func GetIndustries(w http.ResponseWriter, r *http.Request) {
	industries, err := rolesService.GetAllRoles()
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

func CreateRole(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name cannot be empty"))
		return
	}
	err := rolesService.CreateRole(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func EditRole(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	if id == "" || name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id and name cannot be empty"))
		return
	}
	err := rolesService.EditRole(id, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	// Get the request body
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id cannot be empty"))
		return
	}
	err := rolesService.DeleteRole(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
