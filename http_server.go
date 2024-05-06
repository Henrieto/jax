package jax

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HttpServer struct {
	UseTls bool
	Cert   string
	Key    string
	*http.Server
}

func (server *HttpServer) Router(router *Router) {
	server.Handler = router.GetRouter()
}

func (server *HttpServer) LogStartupMsg() {
	start_server_message := "server started : " + server.Addr
	log.Println(start_server_message)
}

func (server *HttpServer) LogShutdownMsg() {
	// log a shut down message
	message := "server is shutting down"
	log.Println(message)
}

func (server *HttpServer) Start() (err error) {
	server.LogStartupMsg()
	if !server.UseTls {
		err = server.ListenAndServe()
	} else {
		if server.Cert == "" && server.Key == "" {
			panic(" server certificate and key is needed ")
		}
		err = server.ListenAndServeTLS(server.Key, server.Cert)
	}
	return
}

func (server *HttpServer) Stop(_context context.Context) (err error) {
	server.LogShutdownMsg()
	// shutdown the server
	err = server.Shutdown(_context)
	return
}

func (server *HttpServer) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, middleware := range middlewares {
		server.Handler = middleware(server.Handler)
	}
}

func (server *HttpServer) Options(options ...ServerOption) {
	for _, option := range options {
		option(server)
	}
}

func (server *HttpServer) Address(address string) {
	// set the server address
	server.Addr = address
}

func (server *HttpServer) Listen() (err error) {
	// start the server in a goroutine
	go func() { err = server.Start() }()
	if err != nil {
		return
	}

	// create a channel for capturing system interrupts
	ch := make(chan os.Signal, 4)

	// attach the system interrupts
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	// make the channel blocking
	<-ch

	// create a timeout context , for the server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// stop the server if there was an interrupt
	server.Stop(ctx)

	return nil
}
