package jax

import (
	"fmt"
	"net/http"

	"github.com/henrieto/jax/command"
)

type Plugin struct {
	Label        string
	Models       map[string]any
	Middlewares  []func(http.Handler) http.Handler
	Routes       []Route
	ModelActions map[string]any
	Commands     []*command.Command
}

func (plugin *Plugin) Attach(router *Router) {
	router.Middlewares = append(router.Middlewares, plugin.Middlewares...)
	if plugin.Label != "" {
		router = router.Prefix(plugin.Label)
	}
	for _, route := range plugin.Routes {
		if route.Method == "" {
			route.Method = http.MethodGet
		}
		if route.Prefix != "" {
			new_router := router.Prefix(route.Prefix)
			new_router.RoutePaths[route.Name] = route.Path
			path := fmt.Sprintf("%v %v", route.Method, router.Path(route.Path))
			new_router.Mux.HandleFunc(path, route.Handler)
			fmt.Println(path)
		} else {
			router.RoutePaths[route.Name] = route.Path
			path := fmt.Sprintf("%v %v", route.Method, router.Path(route.Path))
			router.Mux.HandleFunc(path, route.Handler)
			fmt.Println(path)
		}
	}
}
