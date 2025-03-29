package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

type loggingMiddleware struct {
	logger *slog.Logger
}

// Used for retaining the status code that was written to the [http.ResponseWriter].
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (r *loggingResponseWriter) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

// WithLogging creates a middleware that logs requested routes with the
// requester's IP and the HTTP status code returned.
func WithLogging(logger *slog.Logger) Middleware {
	return &loggingMiddleware{
		logger: logger,
	}
}

func (m *loggingMiddleware) Use(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := newLoggingResponseWriter(w)
		next.ServeHTTP(lw, r)

		m.logger.Info(
			fmt.Sprintf("%s %s", r.Method, r.URL),
			"ip",
			r.RemoteAddr,
			"status_code",
			lw.statusCode,
		)
	})
}
