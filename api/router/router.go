package router

import (
	"net/http"

	"github.com/GotoRen/gacha-api/api/controller"
)

func Router() *http.ServeMux {

	r := http.NewServeMux()
	r.HandleFunc("/api/user/create", controller.CreateUser)
	r.HandleFunc("/api/user/get", controller.GetUser)
	r.HandleFunc("/api/user/login", controller.LoginUser)
	r.HandleFunc("/api/user/update", controller.UpdateUser)
	r.HandleFunc("/api/gacha/draw", controller.DrawGacha)
	r.HandleFunc("/api/character/list", controller.GetUserCharacters)
	r.HandleFunc("/api/check", controller.Check)

	return r
}
