package mangadex

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// MaybeNumber unmarshals numbers that may be numeric, numeric strings or empty
// strings.
type MaybeNumber struct{ json.Number }

// DynamicType represents a field that could be of dynamic type.
type DynamicType struct{ Value interface{} }

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *MaybeNumber) UnmarshalJSON(d []byte) error {
	d = bytes.Trim(d, `"`)
	if string(d) == "" {
		d = []byte("0")
	}
	n.Number = json.Number(d)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *DynamicType) UnmarshalJSON(d []byte) error {
	intValue, err := strconv.ParseInt(string(d), 10, 8)
	if err == nil {
		v.Value = int(intValue)
		return nil
	}

	boolValue, err := strconv.ParseBool(string(d))
	if err == nil {
		v.Value = boolValue
		return nil
	}

	return err
}
