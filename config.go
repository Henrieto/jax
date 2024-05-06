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
	Middlewares  func(http.HandlerFunc) http.HandlerFunc
}
type Config struct {
	Plugins []Plugin
	Router  *RouterConfig
	Server  ServerConfig
}
