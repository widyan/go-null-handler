package null

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type NullTime struct {
	sql.NullTime
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (ns *NullTime) UnmarshalJSON(b []byte) error {
	ns.Valid = string(b) != "null"
	e := json.Unmarshal(b, &ns.Time)
	return e
}
