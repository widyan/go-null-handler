package null

import (
	"database/sql"
	"encoding/json"
)

type NullBool struct {
	sql.NullBool
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON for NullBool
func (ns *NullBool) UnmarshalJSON(b []byte) error {
	ns.Valid = string(b) != "null"
	e := json.Unmarshal(b, &ns.Bool)
	return e
}
