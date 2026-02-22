package null

import (
	"encoding/json"
	"github.com/google/uuid"
)

type NullUuid struct {
	UUID  uuid.UUID
	Valid bool
}

func (nu *NullUuid) Scan(value interface{}) error {
	if value == nil {
		nu.UUID = uuid.UUID{}
		nu.Valid = false
		return nil
	}

	var u uuid.UUID
	err := u.Scan(value)
	if err != nil {
		return err
	}

	nu.UUID = u
	nu.Valid = true
	return nil
}

// MarshalJSON for NullUuid
func (nu *NullUuid) MarshalJSON() ([]byte, error) {
	if !nu.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nu.UUID)
}

// UnmarshalJSON for NullUuid
func (nu *NullUuid) UnmarshalJSON(b []byte) error {
	nu.Valid = string(b) != "null"
	e := json.Unmarshal(b, &nu.UUID)
	return e
}
