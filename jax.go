package jax

import (
	"net/http"

	"github.com/henrieto/jax/command"
)

func New(config *Config) *HttpServer {
	// initalize a new router
	router := &Router{
		Mux:          http.NewServeMux(),
		PrefixString: config.Router.VersionString,
		Middlewares:  []func(http.Handler) http.Handler{},
		RoutePaths:   map[string]string{},
	}
	// attach a prefix to the router if
	// there is a prefix
	if config.Router.Prefix != "" {
		router = router.Prefix(config.Router.Prefix)
	}
	// load the plugins and register there commands
	for _, plugin := range config.Plugins {
		plugin.Attach(router)
		command.Register(plugin.Commands...)
	}
	// set the commands for execution
	command.Execute()
	// initialize a http server
	server := &HttpServer{
		Server: &http.Server{
			WriteTimeout: config.Server.WriteTimeout,
			ReadTimeout:  config.Server.ReadTimeout,
			IdleTimeout:  config.Server.IdleTimeout,
		},
	}
	// attach the router to the server
	server.Router(router)
	// attach the middlewares to the server
	server.Use(router.Middlewares...)
	// set the server port
	server.Address(":" + config.Server.Port)
	// change the default server configuration using options
	for _, option := range config.Server.Options {
		option(server)
	}
	// return the server
	return server
}
