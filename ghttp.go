package main

import (
	"fmt"
	"net/http"
	"time"
)

type GHTTP struct {
	Host        string
	Port        string
	staticDir   string
	server      *http.Server
	multiplexer *http.ServeMux
	middleware  MiddlewareChain
}

func (ghttp *GHTTP) Init() error {
	var err error
	// init middleware
	err = ghttp.InitMiddleware()
	if err != nil {
		return err
	}

	// init handlers
	err = ghttp.InitMultiplexer()
	if err != nil {
		return err
	}
	// init handlers
	ghttp.server = &http.Server{
		Addr:    ghttp.Host + ":" + ghttp.Port,
		Handler: ghttp.multiplexer,
	}

	return nil
}

// starts the server with listenAndServe
func (ghttp *GHTTP) Start() error {
	return ghttp.server.ListenAndServe()
}

func (ghttp *GHTTP) Log(log string) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	prefix := timestamp + ": "
	fmt.Println(prefix + log)
}
