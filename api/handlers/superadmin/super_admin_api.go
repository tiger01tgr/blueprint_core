package superadmin

import (
	"backend/services/admin"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func InitSuperAdminRoutes(router chi.Router) {
	router.Route("/superadmin", func(r chi.Router) {
		r.Post("/login", Login)
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("email and password must be provided"))
		return
	}
	token, err := admin.LoginSuperAdmin(email, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(token))	
}