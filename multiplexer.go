package main

import (
	"net/http"
	"os"
)

// registers all handlers to ghttp.multiplexer

func (ghttp *GHTTP) InitMultiplexer() error {
	ghttp.multiplexer = http.NewServeMux()

	ghttp.multiplexer.Handle("/", ghttp.middleware.Handle(http.HandlerFunc(fileServerHandler)))
	return nil
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	dir := ghttp.staticDir

	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		missingDistHandler(w, r)
		return
	}

	fs := http.FileServer(http.Dir(dir))
	http.StripPrefix("/", fs).ServeHTTP(w, r)
}

func missingDistHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello! no static assets found."))
}
