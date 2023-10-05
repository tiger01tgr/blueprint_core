package middleware

import (
	"backend/config"
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
)

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

		fb := config.GetFirebase()

		var idToken *auth.Token
		var err error
		idToken, err = fb.VerifyIDToken(context.Background(), token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization header provided"))
			return
		}

		ctx := context.WithValue(r.Context(), "email", idToken.Claims["email"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
