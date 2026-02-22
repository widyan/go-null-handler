package null

import (
	"database/sql"
	"encoding/json"
)

type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON for NullFloat64
func (ns *NullFloat64) UnmarshalJSON(b []byte) error {
	ns.Valid = string(b) != "null"
	e := json.Unmarshal(b, &ns.Float64)
	return e
}
