package main

import (
	"io"
	"net/http"
)

// sendErrorResponse 错误信息的响应
func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
