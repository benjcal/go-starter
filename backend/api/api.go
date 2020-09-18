package api

import (
	"github.com/benjcal/go-starter/app"
	"github.com/go-chi/chi"
)

func GetRouter(a app.App) (chi.Router, error) {
	r := chi.NewRouter()

	mountAuthRoutes(r, a)
	mountFunctionRoute(r, a)

	return r, nil
}
