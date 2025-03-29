package ajax

import (
	"encoding/json"
	"time"
)

// Duration is used as a middleman type of [time.Duration] for encoding/decoding
// [time.Duration] to/from JSON.
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(b[1 : len(b)-1]))
	return
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Duration.String())
}
