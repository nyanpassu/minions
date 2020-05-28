package client

import (
	"context"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/projecteru2/minions/lib/ipam"
	log "github.com/sirupsen/logrus"
)

// ReserveIPAddressRequest .
type ReserveIPAddressRequest struct {
	PoolID    string
	IPAddress string
}

// ReleaseAddressRequest .
type ReleaseAddressRequest struct {
	PoolID    string
	IPAddress string
}

// Client .
type Client interface {
	ReserveIPAddress(request *ReserveIPAddressRequest) error
	ReleaseReservedIPAddress(request *ReleaseAddressRequest) error
}

// Options .
type Options struct {
	socketAddress string
}

type client struct {
	etcdv3 *clientv3.Client
	ipam   ipam.IPAddressManager
}

// New .
func New(cfg clientv3.Config) (Client, error) {
	etcdv3, err := clientv3.New(cfg)
	if err != nil {
		return nil, nil
	}
	manager := ipam.New(etcdv3)
	if err != nil {
		return nil, nil
	}
	return client{etcdv3: etcdv3, ipam: manager}, nil
}

func (cli client) ReserveIPAddress(request *ReserveIPAddressRequest) error {
	return cli.ipam.ReserveIPAddress(context.Background(), request.PoolID, request.IPAddress)
}

func (cli client) ReleaseReservedIPAddress(request *ReleaseAddressRequest) error {
	ctx := context.Background()
	released, err := cli.ipam.ReleaseReservedIPAddress(ctx, request.PoolID, request.IPAddress)
	if err != nil {
		return err
	}
	if !released {
		err = errors.Errorf("reserved ip has already been reallocated, ip: %s", request.IPAddress)
		log.Errorln(err)
		return err
	}
	if err := cli.ipam.ReleaseCalicoIPAddress(ctx, request.PoolID, request.IPAddress); err != nil {
		return err
	}
	return nil
}
