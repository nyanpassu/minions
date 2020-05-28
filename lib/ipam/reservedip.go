package ipam

import (
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/mvcc/mvccpb"
)

// ReservedIPAddress .
type ReservedIPAddress struct {
	PoolID    string
	IPAddress string
	CreateAt  int64
	Revision  int64
}

func pathToPool(poolID string) string {
	return fmt.Sprintf("/barrel/reserved/pool/%s/", poolID)
}

func pathToIPAddress(poolID string, ipAddress string) string {
	return fmt.Sprintf("/barrel/reserved/pool/%s/ip/%s", poolID, ipAddress)
}

func kvToModel(ekv *mvccpb.KeyValue) (ReservedIPAddress, error) {
	result := ReservedIPAddress{}
	if err := json.Unmarshal(ekv.Value, &result); err != nil {
		return result, err
	}
	return result, nil
}

func modelToKv(model ReservedIPAddress) (string, string) {
	key := pathToIPAddress(model.PoolID, model.IPAddress)
	// adopt json.Marshal will introduce handling error, so use json template
	return key, `{"PoolID":"%s", "IPAddress":"%s", "CreateAt":%d}`
}
