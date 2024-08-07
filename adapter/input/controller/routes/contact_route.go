package routes

import (
	"challenger/adapter/input/controller"
	"net/http"
)

func ContactRouter(cc controller.ControllerInterface) {
	http.HandleFunc("GET /", cc.Form)
	http.HandleFunc("POST /page", cc.CreateContact)
}
