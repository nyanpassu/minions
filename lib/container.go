package lib

import (
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/mvcc/mvccpb"
)

// Container .
type Container struct {
	ID      string
	PoolID  string
	Address string
	version int64
}

// Key .
func (container *Container) Key() string {
	if container.ID == "" {
		return ""
	}
	return fmt.Sprintf("/barrel/containers/%s", container.ID)
}

// Read .
func (container *Container) Read(ekv *mvccpb.KeyValue) error {
	container.version = ekv.Version
	return json.Unmarshal(ekv.Value, container)
}

// JSON .
func (container *Container) JSON() ([]byte, error) {
	return json.Marshal(container)
}

// Version .
func (container *Container) Version() int64 {
	return container.version
}
