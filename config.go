package jax

import (
	"net/http"
	"time"
)

type RouterConfig struct {
	VersionString string
	Version       int
	Prefix        string
}

type ServerConfig struct {
	Port         string
	Options      []ServerOption
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
	Middlewares  []func(http.HandlerFunc) http.HandlerFunc
}
type Config struct {
	Plugins []Plugin
	Router  *RouterConfig
	Server  *ServerConfig
}

func DefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:         "8080",
		Options:      []ServerOption{},
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 30,
		Middlewares:  []func(http.HandlerFunc) http.HandlerFunc{},
	}
}
