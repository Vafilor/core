package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONLabels map[string]string

// Unmarshal unmarshal's the json in j to v, as in json.Unmarshal.
func (l *JSONLabels) Unmarshal(v interface{}) error {
	if len(*l) == 0 {
		*l = make(map[string]string)
	}
	v = l
	return nil
}

// Value returns j as a value.  This does a validating unmarshal into another
// RawMessage.  If j is invalid json, it returns an error.
func (l JSONLabels) Value() (driver.Value, error) {
	return json.Marshal(l)
}

// Scan stores the src in *j.  No validation is done.
func (l *JSONLabels) Scan(src interface{}) error {
	var source []byte
	switch t := src.(type) {
	case string:
		source = []byte(t)
	case []byte:
		if len(t) == 0 {
			source = []byte("{}")
		} else {
			source = t
		}
	case nil:
		*l = make(map[string]string)
	default:
		return errors.New("incompatible type for JSONLabels")
	}

	return json.Unmarshal(source, l)
}
