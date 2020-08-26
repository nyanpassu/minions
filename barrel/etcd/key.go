// package etcd

// import (
// 	"fmt"

// 	"github.com/projecteru2/minions/types"
// )

// const (
// 	barrelAddressKeyPrefx             = "/barrel/addresses/%s"
// 	barrelPoolAddressKeyPrefx         = "/barrel/pools/%s/addresses/%s"
// 	barrelContainersKeyPrefx          = "/barrel/containers/%s"
// 	barrelReserverRequestKeyPrefx     = "/barrel/reservereqs/%s"
// 	barrelPoolReserverRequestKeyPrefx = "/barrel/pools/%s/reservereqs/%s"
// )

// func keyOfReservedAddress(address *types.ReservedAddress) string {
// 	if address.PoolID == "" {
// 		return fmt.Sprintf(barrelAddressKeyPrefx, address.Address)
// 	}
// 	return fmt.Sprintf(barrelPoolAddressKeyPrefx, address.PoolID, address.Address)
// }

// func keyOfContainerInfo(info *types.ContainerInfo) string {
// 	return fmt.Sprintf(barrelContainersKeyPrefx, info.ID)
// }

// func keyOfReserveRequest(req *types.ReserveRequest) string {
// 	if req.PoolID == "" {
// 		return fmt.Sprintf(barrelReserverRequestKeyPrefx, req.Address)
// 	}
// 	return fmt.Sprintf(barrelPoolReserverRequestKeyPrefx, req.PoolID, req.Address)
// }