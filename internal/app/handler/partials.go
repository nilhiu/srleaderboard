package handler

import (
	"context"
	"net/http"

	"github.com/nilhiu/srleaderboard/internal/app/htmx"
	"github.com/nilhiu/srleaderboard/internal/app/middleware"
	"github.com/nilhiu/srleaderboard/internal/service/user"
	"github.com/nilhiu/srleaderboard/internal/view/component"
)

// PartialsHandler provides methods to handle HTMX requests for partial HTMLs.
type PartialsHandler struct {
	ctx context.Context
}

// NewPartialsHandler creates a new PartialsHandler
func NewPartialsHandler(ctx context.Context) *PartialsHandler {
	return &PartialsHandler{
		ctx: ctx,
	}
}

// Navbar is a handler, which returns partial HTML for the navbar. Only supports
// HTMX requests (`Hx-Request: "true"`)
func (h *PartialsHandler) Navbar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Hx-Request") != "true" {
			w.WriteHeader(http.StatusForbidden)
		}

		if tokStr := r.Context().Value(middleware.ContextValueKeyJWT); tokStr != nil {
			_, err := user.ValidateJWT(tokStr.(string))
			if err == nil {
				htmx.MustRender(h.ctx, w, component.Navbar(true))
				return
			}
		}

		htmx.MustRender(h.ctx, w, component.Navbar(false))
	}
}
