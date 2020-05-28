package ipam

import (
	"context"

	"github.com/coreos/etcd/clientv3"
)

// IPAddressManager .
type IPAddressManager interface {
	GetReservedIPAddress(ctx context.Context, poolID string, ipAddress string) (ReservedIPAddress, bool, error)
	ListReservedIPAddress(ctx context.Context, poolID string) ([]ReservedIPAddress, error)
	ReserveIPAddress(ctx context.Context, poolID string, ipAddress string) error
	ReleaseReservedIPAddress(ctx context.Context, poolID string, ipAddress string) (bool, error)
	ReleaseCalicoIPAddress(ctx context.Context, poolID string, ipAddress string) error
}

type ipAddrManager struct {
	etcdV3 *clientv3.Client
}

// New .
func New(etcdV3 *clientv3.Client) IPAddressManager {
	return ipAddrManager{etcdV3: etcdV3}
}

func (ipam ipAddrManager) GetReservedIPAddress(ctx context.Context, poolID string, ipAddress string) (ReservedIPAddress, bool, error) {
	key := pathToIPAddress(poolID, ipAddress)
	resp, err := ipam.etcdV3.Get(ctx, key)
	if err != nil {
		return ReservedIPAddress{}, false, err
	}
	if len(resp.Kvs) == 0 {
		return ReservedIPAddress{}, false, nil
	}
	model, err := kvToModel(resp.Kvs[0])
	if err != nil {
		return ReservedIPAddress{}, false, err
	}
	model.Revision = resp.Header.Revision
	return model, true, nil
}

// ListReservedIPAddress .
func (ipam ipAddrManager) ListReservedIPAddress(ctx context.Context, poolID string) ([]ReservedIPAddress, error) {
	// TODO
	return nil, nil
}

// ReserveIPAddress .
func (ipam ipAddrManager) ReserveIPAddress(ctx context.Context, poolID string, ipAddress string) error {
	model := ReservedIPAddress{
		PoolID:    poolID,
		IPAddress: ipAddress,
	}
	key, value := modelToKv(model)
	_, err := ipam.etcdV3.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

// ReleaseIPAddress .
func (ipam ipAddrManager) ReleaseReservedIPAddress(ctx context.Context, poolID string, ipAddress string) (bool, error) {
	key := pathToIPAddress(poolID, ipAddress)
	resp, err := ipam.etcdV3.Delete(ctx, key)
	if err != nil {
		return false, err
	}
	return resp.Deleted > 0, nil
}

// ReleaseCalicoIPAddress .
func (ipam ipAddrManager) ReleaseCalicoIPAddress(ctx context.Context, poolID string, ipAddress string) error {
	// TODO
	return nil
}
