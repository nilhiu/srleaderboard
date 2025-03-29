package middleware

import (
	"context"
	"net/http"

	"github.com/nilhiu/srleaderboard/internal/service/user"
)

type protectedMiddleware struct{}

// WithProtected creates a middleware which requires that the request context
// has the value for [ContextValueKeyJWT]. If it doesn't it returns the status
// code of `Unauthorized` and skips calling the `next` handler.
func WithProtected() Middleware {
	return &protectedMiddleware{}
}

func (m *protectedMiddleware) Use(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokStr, ok := r.Context().Value(ContextValueKeyJWT).(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tok, err := user.ValidateJWT(tokStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		sub, err := tok.Claims.GetSubject()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), ContextValueKeyUser, sub)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
