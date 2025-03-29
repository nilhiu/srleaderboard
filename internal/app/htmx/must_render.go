package htmx

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

// MustRender renders the Templ component to the [http.ResponseWriter] ignoring
// the error.
func MustRender(ctx context.Context, w http.ResponseWriter, c templ.Component) {
	w.Header().Add("Content-Type", "text/html")

	if err := c.Render(ctx, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
