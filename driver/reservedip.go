package driver

import (
	log "github.com/sirupsen/logrus"

	"github.com/coreos/etcd/clientv3"
	"github.com/projecteru2/minions/lib"
	"github.com/projecteru2/minions/utils"
)

const fixedIPLabel = "fixed-ip"

type reservedIPManager struct {
	etcd utils.EtcdClient
}

// ReservedIPManager .
type ReservedIPManager interface {
	Reserve(address lib.ReservedIPAddress, containerID string) error
	AquireIfReserved(address lib.ReservedIPAddress) (bool, error)
	IsReserved(address lib.ReservedIPAddress) (bool, error)
	// marked is not guaranteed to be false when err is not nil
	ConsumeRequestMarkIfPresent(request lib.ReserveRequest) (marked bool, err error)
}

// NewReservedIPManager .
func NewReservedIPManager(etcdV3Client *clientv3.Client) ReservedIPManager {
	return reservedIPManager{etcd: utils.EtcdClient{Etcd: etcdV3Client}}
}

// Reserve .
func (ripam reservedIPManager) Reserve(ip lib.ReservedIPAddress, containerID string) error {
	log.Infof("Reserve ip = %v", ip)
	return ripam.etcd.PutMulti(&ip, &lib.Container{ID: containerID, PoolID: ip.PoolID, Address: ip.Address})
}

// IsReserved .
func (ripam reservedIPManager) IsReserved(ip lib.ReservedIPAddress) (bool, error) {
	return ripam.etcd.Get(&ip)
}

// IsRequested .
func (ripam reservedIPManager) ConsumeRequestMarkIfPresent(request lib.ReserveRequest) (bool, error) {
	return ripam.etcd.GetAndDelete(&request)
}

// AquireIfReserved .
func (ripam reservedIPManager) AquireIfReserved(ip lib.ReservedIPAddress) (bool, error) {
	log.Infof("AquireIPIfReserved ip = %v", ip)
	return ripam.etcd.GetAndDelete(&ip)
}
