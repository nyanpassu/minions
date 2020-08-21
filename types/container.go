package types

import "encoding/json"

// Container .
type Container struct {
	Base
	ID string
}

// JSON .
func (c *Container) JSON() ([]byte, error) {
	return json.Marshal(c)
}
