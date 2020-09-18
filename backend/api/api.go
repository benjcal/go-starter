package api

import (
	"github.com/benjcal/go-starter/app"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func GetRouter(a app.App) (chi.Router, error) {
	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)

	mountAuthRoutes(r, a)
	mountFunctionRoute(r, a)

	return r, nil
}
