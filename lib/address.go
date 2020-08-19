package lib

import (
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/mvcc/mvccpb"
)

// ReservedIPAddress .
type ReservedIPAddress struct {
	PoolID  string
	Address string
	version int64
}

// Key .
func (addr *ReservedIPAddress) Key() string {
	if addr.Address == "" {
		return ""
	}
	if addr.PoolID == "" {
		return fmt.Sprintf("/barrel/addresses/%s", addr.Address)
	}
	return fmt.Sprintf("/barrel/pools/%s/addresses/%s", addr.PoolID, addr.Address)
}

// Read .
func (addr *ReservedIPAddress) Read(ekv *mvccpb.KeyValue) error {
	addr.version = ekv.Version
	return json.Unmarshal(ekv.Value, addr)
}

// JSON .
func (addr *ReservedIPAddress) JSON() ([]byte, error) {
	return json.Marshal(addr)
}

// Version .
func (addr *ReservedIPAddress) Version() int64 {
	return addr.version
}
