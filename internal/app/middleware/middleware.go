package middleware

import "net/http"

// Middleware defines an interface which is used to add additional processing
// to HTTP requests. It's implementations should use the `Use` method to wrap
// the given handler and return a new one that implements the additional processing.
type Middleware interface {
	// Use applies the middleware to the given [http.Handler].
	Use(next http.Handler) http.Handler
}
