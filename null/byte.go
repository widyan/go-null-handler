package null

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type NullBytes struct {
	Bytes []byte
	Valid bool // Valid is true if Bytes is not NULL
}

// Scan implements the Scanner interface.
func (nb *NullBytes) Scan(value interface{}) error {
	if value == nil {
		nb.Bytes, nb.Valid = nil, false
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("NullBytes: type assertion to []byte failed")
	}
	nb.Bytes = b
	nb.Valid = true
	return nil
}

// Value implements the driver Valuer interface.
func (nb NullBytes) Value() (driver.Value, error) {
	if !nb.Valid {
		return nil, nil
	}
	return nb.Bytes, nil
}

func (nb NullBytes) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bytes) // will marshal as a base64-encoded string (default for []byte)
}

func (nb *NullBytes) UnmarshalJSON(data []byte) error {
	// Handle null
	if string(data) == "null" {
		nb.Bytes = nil
		nb.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &nb.Bytes); err != nil {
		return err
	}
	nb.Valid = true
	return nil
}
