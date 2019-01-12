package http

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

// WithBackend provides an option to set the http backend url
func WithBackend(url string) micro.Option {
	return func(o *micro.Options) {
		// get the router
		r := o.Server.Options().Router
		// check its a http router
		if httpRouter, ok := r.(*Router); ok {
			httpRouter.Backend = url
		}
	}
}

// WithRouter provides an option to set the http router
func WithRouter(r server.Router) micro.Option {
	return func(o *micro.Options) {
		o.Server.Init(server.WithRouter(r))
	}
}
