package jax

import (
	"net/http"
	"time"
)

type Config struct {
	Plugins []Plugin
	Router  struct {
		VersionString string
		Version       int
		Prefix        string
	}
	Server struct {
		Port         string
		Options      []ServerOption
		WriteTimeout time.Duration
		ReadTimeout  time.Duration
		IdleTimeout  time.Duration
		Middlewares  func(http.HandlerFunc) http.HandlerFunc
	}
}
