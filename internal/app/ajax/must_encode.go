package ajax

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// MustEncode writes the JSON encoding of `v` to `w`, but doesn't return an error.
func MustEncode(w http.ResponseWriter, v any) {
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		slog.Error("must_encode", "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
