package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Data map[string]interface{}

type Header struct {
	ID      int64  `json:"id" db:"id"`
	Version string `json:"version" db:"version"`
	Attrs   Data   `json:"data" db:"data"`
}

func (a Data) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Data) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
