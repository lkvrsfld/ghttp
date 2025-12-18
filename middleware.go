package main

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler
type MiddlewareChain []Middleware

// registers all middleware and writes to ghttp.middlewares
func (ghttp *GHTTP) InitMiddleware() error {
	ghttp.middleware = append(ghttp.middleware, logMiddleware)
	return nil
}

// main handle function, iterates through all middleware from first to last, and handles the route
func (mc MiddlewareChain) Handle(originalHandler http.Handler) http.Handler {
	if originalHandler == nil {
		originalHandler = http.DefaultServeMux
	}

	for i := range mc {
		originalHandler = mc[len(mc)-1-i](originalHandler)
	}
	return originalHandler
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := r.Method + " " + r.URL.Path + " from " + r.RemoteAddr
		ghttp.Log(msg)
		h.ServeHTTP(w, r)
	})
}
