package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//  middleWareHandler 路由中间件
type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

// NewMiddleWareHandler 初始化路由中间件
func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

// RegisterHandlers 实现注册路由
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", streamHandler)

	router.POST("/upload/:vid-id", uploadHandler)

	router.GET("/testpage", testPageHandler)

	return router
}

// ServeHTTP 使路由器实现http.Handler接口
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}
