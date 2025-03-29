package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nilhiu/srleaderboard/internal/app/htmx"
	"github.com/nilhiu/srleaderboard/internal/app/middleware"
	"github.com/nilhiu/srleaderboard/internal/view/page"
)

// UserPageHandler provides methods to handle HTMX requests to get the user page.
type UserPageHandler struct {
	ctx context.Context
}

func NewUserPageHandler(ctx context.Context) *UserPageHandler {
	return &UserPageHandler{
		ctx: ctx,
	}
}

// User is a handler, which returns HTML of the UserPage component. The specific
// user is provided as a path value `user`.
func (h *UserPageHandler) User() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.PathValue("user")
		htmx.MustRender(h.ctx, w, page.UserPage(user))
	}
}

// Profile is a handler, which returns HTML of the UserPage component for the
// current authenticated user.
func (h *UserPageHandler) Profile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middleware.ContextValueKeyUser).(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if r.Header.Get("Hx-Request") == "true" {
			w.Header().Add("Location", fmt.Sprintf("/api/runs/%s?offset=0&limit=5", user))
			w.WriteHeader(http.StatusFound)
			return
		}

		htmx.MustRender(h.ctx, w, page.UserPage(user))
	}
}
