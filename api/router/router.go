package router

import (
	"net/http"

	"github.com/fumist23/game-api/controller"
)

func Router() *http.ServeMux {

	r := http.NewServeMux()
	r.HandleFunc("/user/create", controller.CreateUser)
	r.HandleFunc("/user/get", controller.GetUser)
	r.HandleFunc("/user/login", controller.LoginUser)
	r.HandleFunc("/user/update", controller.UpdateUser)
	r.HandleFunc("/gacha/draw", controller.DrawGacha)
	r.HandleFunc("/character/list", controller.GetUserCharacters)
	r.HandleFunc("/check", controller.Check)

	return r
}
