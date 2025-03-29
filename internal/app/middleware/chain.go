package middleware

import "net/http"

// A Chain represents a list of middleware, which can be applied to a handler
// in succession.
type Chain struct {
	middlewares []Middleware
}

// NewChain creates a new chain with the given middleware
func NewChain(middlewares ...Middleware) *Chain {
	return &Chain{
		middlewares: middlewares,
	}
}

// Use applies the middleware in the chain to the given [http.Handler].
func (c *Chain) Use(next http.Handler) http.Handler {
	h := next
	for _, m := range c.middlewares {
		h = m.Use(h)
	}

	return h
}
