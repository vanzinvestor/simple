package middleware

import "github.com/julienschmidt/httprouter"

type middleware func(httprouter.Handle) httprouter.Handle

// Chain middleware chaining middleware
func Chain(ms ...middleware) middleware {
	return func(next httprouter.Handle) httprouter.Handle {
		for i := len(ms); i > 0; i-- {
			next = ms[i-1](next)
		}
		return next
	}
}
