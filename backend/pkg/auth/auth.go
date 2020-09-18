package auth

import (
	"github.com/benjcal/go-starter/app/models"
	"github.com/go-chi/jwtauth"
)

func GetJWTAuth(c models.Config) *jwtauth.JWTAuth {
	a := jwtauth.New("HS256", []byte(c.JWTSecret), nil)

	return a
}
