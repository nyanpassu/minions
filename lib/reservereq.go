package lib

import (
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/mvcc/mvccpb"
)

// ReserveRequest .
type ReserveRequest struct {
	PoolID  string
	Address string
	version int64
}

// Key .
func (req *ReserveRequest) Key() string {
	if req.Address == "" {
		return ""
	}
	if req.PoolID == "" {
		return fmt.Sprintf("/barrel/reservereqs/%s", req.Address)
	}
	return fmt.Sprintf("/barrel/pools/%s/reservereqs/%s", req.PoolID, req.Address)
}

// Read .
func (req *ReserveRequest) Read(ekv *mvccpb.KeyValue) error {
	req.version = ekv.Version
	return json.Unmarshal(ekv.Value, req)
}

// JSON .
func (req *ReserveRequest) JSON() ([]byte, error) {
	return json.Marshal(req)
}

// Version .
func (req *ReserveRequest) Version() int64 {
	return req.version
}
