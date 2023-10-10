package jwt

import (
	"sync"

	"github.com/go-chi/jwtauth/v5"
)

var (
	jwt *jwtauth.JWTAuth
	once sync.Once
)

func GetJWT() *jwtauth.JWTAuth {
	once.Do(func() {
		jwt = jwtauth.New("HS256", []byte("secret"), nil)
	})
	return jwt
}