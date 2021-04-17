package main

import (
	"log"
	"net/http"

	"github.com/GotoRen/gacha-api/api/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// CORS対策
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	r := router.Router()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
