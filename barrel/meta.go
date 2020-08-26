// package barrel

// import (
// 	"context"

// 	"github.com/projecteru2/minions/types"
// )

// // MetaInterface .
// type MetaInterface interface {
// 	ReserveIPforContainer(ctx context.Context, IP *types.ReservedAddress, ID string) error
// 	IPIsReserved(ctx context.Context, IP *types.ReservedAddress) (bool, error)
// 	ConsumeRequestMarkIfPresent(ctx context.Context, req *types.ReserveRequest) (bool, error)
// 	AquireIfReserved(ctx context.Context, IP *types.ReservedAddress) (bool, error)
// }
