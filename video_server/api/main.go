package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers 用户注册路由
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}
