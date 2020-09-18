package main

import (
	"fmt"
	"github.com/benjcal/go-starter/api"
	"github.com/benjcal/go-starter/app"
	"github.com/benjcal/go-starter/pkg/conf"
	"log"
	"net/http"
)

func main() {
	c, err := conf.GetConfig(`db_dsn = "user=postgres password=postgres dbname=go port=5432 sslmode=disable"
jwt_secret = "qwerty"`)
	if err != nil {
		log.Fatal(err)
	}

	a, err := app.Init(c)
	if err != nil {
		log.Fatal(err)
	}

	r, err := api.GetRouter(a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("listening on port 3000...")
	http.ListenAndServe(":3000", r)
}
