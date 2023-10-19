package middleware

import (
	"backend/config/firebase"
	"context"
	"net/http"
	"strings"
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"backend/services/jwt"
	"backend/services/user"
)

func SuperAdminAuth(router chi.Router) {
	router.Use(jwtauth.Verifier(jwt.GetJWT()))
	router.Use(jwtauth.Authenticator)
}

func UserAuth(router chi.Router) {
	router.Use(GoogleAuth)
}

func GoogleAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// No authorization header provided
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("No authorization header provided"))
			return
		}

		// The header value is expected to be of the format `Bearer {token-body}`,
		// so we'll split on the space and grab the token body.
		token := strings.Split(authHeader, " ")[1]

		fb := firebase.GetFirebase()
		var idToken *auth.Token
		var err error
		idToken, err = fb.VerifyIDToken(context.Background(), token)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization header provided"))
			return
		}
		email := idToken.Claims["email"]
		id, err := user.GetUserIdWithEmail(email.(string))
		if err != nil {
			if err.Error() == "No user found for email" {
				// Create user
				user.CreateUser("", "", "", email.(string), "user")
				id, err = user.GetUserIdWithEmail(email.(string))
				if err != nil {
					fmt.Println(err)
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Invalid authorization header provided"))
					return
				}
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid authorization header provided"))
				return
			}
		}
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}