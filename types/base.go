package types

import "encoding/json"

// Base .
type Base struct {
	PoolID  string
	Address string
}

// JSON .
func (b *Base) JSON() ([]byte, error) {
	return json.Marshal(b)
}
