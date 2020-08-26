// package etcd

// import (
// 	"context"
// 	"fmt"

// 	"github.com/etcd-io/etcd/clientv3"
// 	"github.com/projecteru2/minions/types"
// )

// // ReserveIPforContainer .
// func (e *Etcd) ReserveIPforContainer(ctx context.Context, IP *types.ReservedAddress, ID string) error {
// 	container := &types.ContainerInfo{
// 		ID: ID,
// 		ReservedAddress: types.ReservedAddress{
// 			PoolID:  IP.PoolID,
// 			Address: IP.Address,
// 		},
// 	}
// 	ipKey := fmt.Sprintf(barrelAddressKeyPrefx, IP.Address)
// 	if IP.PoolID != "" {
// 		ipKey = fmt.Sprintf(barrelPoolAddressKeyPrefx, IP.PoolID, IP.Address)
// 	}
// 	ipValue, err := IP.JSON()
// 	if err != nil {
// 		return err
// 	}
// 	containerKey := keyOfContainerInfo(container)
// 	containerValue, err := container.JSON()
// 	if err != nil {
// 		return err
// 	}

// 	data := map[string]string{
// 		ipKey:        string(ipValue),
// 		containerKey: string(containerValue),
// 	}
// 	_, err = e.BatchPut(ctx, data, nil)
// 	return err
// }

// // IPIsReserved .
// func (e *Etcd) IPIsReserved(ctx context.Context, IP *types.ReservedAddress) (bool, error) {
// 	ipKey := fmt.Sprintf(barrelAddressKeyPrefx, IP.Address)
// 	if IP.PoolID != "" {
// 		ipKey = fmt.Sprintf(barrelPoolAddressKeyPrefx, IP.PoolID, IP.Address)
// 	}
// 	resp, err := e.Get(ctx, ipKey)
// 	if err != nil {
// 		return false, err
// 	}
// 	return len(resp.Kvs) > 0, nil
// }

// // ConsumeRequestMarkIfPresent .
// func (e *Etcd) ConsumeRequestMarkIfPresent(ctx context.Context, req *types.ReserveRequest) (bool, error) {
// 	reqKey := fmt.Sprintf(barrelReserverRequestKeyPrefx, req.Address)
// 	if req.PoolID != "" {
// 		reqKey = fmt.Sprintf(barrelPoolReserverRequestKeyPrefx, req.PoolID, req.Address)
// 	}

// 	resp, err := e.Delete(ctx, reqKey, clientv3.WithPrevKV())
// 	if err != nil {
// 		return false, err
// 	}

// 	return len(resp.PrevKvs) > 0, nil
// }

// // AquireIfReserved .
// func (e *Etcd) AquireIfReserved(ctx context.Context, IP *types.ReservedAddress) (bool, error) {
// 	ipKey := fmt.Sprintf(barrelAddressKeyPrefx, IP.Address)
// 	if IP.PoolID != "" {
// 		ipKey = fmt.Sprintf(barrelPoolAddressKeyPrefx, IP.PoolID, IP.Address)
// 	}

// 	resp, err := e.Delete(ctx, ipKey, clientv3.WithPrevKV())
// 	if err != nil {
// 		return false, err
// 	}

// 	return len(resp.PrevKvs) > 0, nil
// }
