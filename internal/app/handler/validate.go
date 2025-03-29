package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/nilhiu/srleaderboard/internal/app/ajax"
	"github.com/nilhiu/srleaderboard/internal/app/htmx"
	"github.com/nilhiu/srleaderboard/internal/view/component"
)

// ValidateHandler provides methods to validate HTML input fields with the help
// of HTMX.
type ValidateHandler struct {
	ctx context.Context
}

func NewValidateHandler(ctx context.Context) *ValidateHandler {
	return &ValidateHandler{
		ctx: ctx,
	}
}

// Time is a handler, which validates if the given `time` for a run is in the
// correct format. The request must be done via HTMX, or `HX-Request: "true"`
// header field has to exist.
func (h *ValidateHandler) Time() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "true" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		var req ajax.ValidateTimeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p := component.ValidatedInputProps{
			ID:             "run-time-input",
			Input:          "text",
			Name:           "time",
			LabelText:      "Time",
			ValidatorRoute: "/api/validate/time",
			Validity:       true,
			Value:          req.Time,
		}

		w.Header().Add("Content-Type", "text/html")

		_, err := time.ParseDuration(req.Time)
		if err != nil {
			p.Validity = false
			p.InvalidMessage = "Incorrect format, please make sure it's the same as shown above"
			htmx.MustRender(h.ctx, w, component.ValidatedInput(p))
		} else {
			htmx.MustRender(h.ctx, w, component.ValidatedInput(p))
		}
	}
}
