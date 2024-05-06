package jax

import (
	"net/http"

	"github.com/henrieto/jax/command"
)

type Jax struct {
	Server  *HttpServer
	Plugins []Plugin
	Router  *Router
	Options []ServerOption
}

func (jx *Jax) AttachPlugins() {
	for _, plugin := range jx.Plugins {
		plugin.Attach(jx.Router)
		command.Register(plugin.Commands...)
	}
}

func (jx *Jax) Initialize() {
	jx.AttachPlugins()
	// set the commands for execution
	command.Execute()
	// change the default server configuration using options
	for _, option := range jx.Options {
		option(jx.Server)
	}
}

func New(config *Config) *Jax {
	// initalize a new router
	if config.Router == nil {
		config.Router = &RouterConfig{}
	}

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
	// initalize a http server
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

	jx := &Jax{
		Server:  server,
		Plugins: config.Plugins,
		Router:  router,
		Options: config.Server.Options,
	}
	// return the Jax object
	return jx
}
