package api

import (
	"encoding/json"
	"github.com/benjcal/go-starter/app"
	"github.com/benjcal/go-starter/pkg/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func mountAuthRoutes(r chi.Router, a app.App) {
	r.Post("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		body, err := getMapFromReqBody(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		v, err := a.UserLogin(body["email"].(string), body["pwd"].(string))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := json.Marshal(v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(res))
	})

	r.Post("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		body, err := getMapFromReqBody(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		err = a.UserRegister(body["email"].(string), body["pwd"].(string), body["name"].(string))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

}

func mountFunctionRoute(r chi.Router, a app.App) {
	jwtAuth := auth.GetJWTAuth(a.Config)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

	})
}

func getMapFromReqBody(r *http.Request) (v map[string]interface{}, err error) {

	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
