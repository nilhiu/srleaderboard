package middleware

import (
	"context"
	"net/http"
)

type jwtMiddleware struct{}

// WithJWT creates a middleware that fetches the JWT from the client's cookies.
func WithJWT() Middleware {
	return &jwtMiddleware{}
}

func (m *jwtMiddleware) Use(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), ContextValueKeyJWT, c.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
