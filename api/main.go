package main

import (
	"log"
	"net/http"

	"github.com/GotoRen/gacha-api/api/router"
)

func main() {

	r := router.Router()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
