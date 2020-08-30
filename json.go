package mangadex

import (
	"bytes"
	"encoding/json"
)

// MaybeNumber unmarshals numbers that may be numeric, numeric strings or empty
// strings.
type MaybeNumber struct{ json.Number }

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *MaybeNumber) UnmarshalJSON(d []byte) error {
	d = bytes.Trim(d, `"`)
	if string(d) == "" {
		d = []byte("0")
	}
	n.Number = json.Number(d)
	return nil
}
