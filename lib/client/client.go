package client

import "github.com/docker/go-plugins-helpers/ipam"

// ReserveIPAddressRequest .
type ReserveIPAddressRequest struct {
	PoolID  string
	Address string
}

// Client .
type Client interface {
	ReserveIPAddress(request *ReserveIPAddressRequest) error
	ReleaseReservedIPAddress(request *ipam.ReleaseAddressRequest) error
}

// New .
func New() Client {
	// TODO
	return nil
}
