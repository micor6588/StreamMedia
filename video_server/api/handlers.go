package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateUser 创建用户信息处理函数
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

// Login 创建用户信息处理函数
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name")
	io.WriteString(w, userName)
}
