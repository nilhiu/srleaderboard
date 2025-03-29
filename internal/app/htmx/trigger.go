package htmx

import (
	"encoding/json"
	"net/http"
)

// A Trigger represents the HTMX's `HX-Trigger` header field.
type Trigger struct {
	events map[string]map[string]string
}

// NewTrigger returns an empty Trigger.
func NewTrigger() *Trigger {
	t := Trigger{
		events: map[string]map[string]string{},
	}

	return &t
}

// AlertOK places a trigger for the frontend to show an successful alert.
func (t *Trigger) AlertOK(msg string) *Trigger {
	t.Add("show-alert-ok",
		"target", "#alert-ok",
		"message", msg)
	return t
}

// AlertOK places a trigger for the frontend to show an erroneous alert.
func (t *Trigger) AlertError(msg string) *Trigger {
	t.Add("show-alert-error",
		"target", "#alert-error",
		"message", msg)
	return t
}

// Add adds a event to trigger with the given props, which have to be placed
// as key-value pairs. i.e. `Add("event", "key", "value", ...)`.
func (t *Trigger) Add(event string, props ...string) *Trigger {
	t.addEvent(event, props)
	return t
}

// Write writes the accumulated events to the [http.ResponseWriter]'s `HX-Trigger`
// header field.
func (t *Trigger) Write(w http.ResponseWriter) {
	j, err := json.Marshal(t.events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", string(j))
}

// Helper function for adding events to a trigger.
func (t *Trigger) addEvent(event string, props []string) {
	if len(props) == 0 {
		t.events[event] = map[string]string{}
	}

	ps := map[string]string{}
	for i := 0; i < len(props); i += 2 {
		ps[props[i]] = props[i+1]
	}

	t.events[event] = ps
}
