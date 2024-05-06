package jax

import (
	"fmt"
	"net/http"
)

type Route struct {
	Prefix  string
	Path    string
	Handler http.HandlerFunc
	Method  string
	Name    string
}

type Router struct {
	Mux           *http.ServeMux
	PrefixString  string
	VersionString string
	Middlewares   []func(http.Handler) http.Handler
	RoutePaths    map[string]string
}

func (router *Router) Path(path string) string {
	if router.VersionString == "" {
		path = fmt.Sprintf("%v%v", router.PrefixString, path)
		return path
	}
	path = fmt.Sprintf("%v%v%v", router.VersionString, router.PrefixString, path)
	return path
}

func (router *Router) GetRouter() *http.ServeMux {
	return router.Mux
}

// Version() --> give unique id to different versions of the same path
// e.g
//  1. api/1/users/create
//  2. api/2/users/create
func (router *Router) Version(version int) *Router {
	if router.VersionString == "" {
		router.VersionString = "/api"
	}
	if version <= 0 {
		version = 1
	}
	router.VersionString = fmt.Sprintf("%v/%v", router.VersionString, version)
	return router
}

// Prefix() --> add prefix to the url path
func (router *Router) Prefix(prefix string) *Router {
	new_router := &Router{
		Mux:           router.Mux,
		PrefixString:  router.PrefixString,
		VersionString: router.VersionString,
		Middlewares:   router.Middlewares,
		RoutePaths:    router.RoutePaths,
	}
	new_router.PrefixString = fmt.Sprintf("%v%v", router.PrefixString, prefix)
	return new_router
}
