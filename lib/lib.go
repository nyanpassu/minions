package lib

import (
	"fmt"

	"github.com/coreos/etcd/clientv3"
	libcalico "github.com/projectcalico/libcalico-go/lib/clientv3"
	caliconet "github.com/projectcalico/libcalico-go/lib/net"
	"github.com/projecteru2/minions/calico"
	"github.com/projecteru2/minions/utils"
	log "github.com/sirupsen/logrus"
)

// Client .
type Client interface {
	RequestFixedIP(poolID string) (string, error)
	MarkFixedIPForContainer(containerID string, address ReservedIPAddress) error
	ReleaseReservedIPByTiedContainerIDIfIdle(containerID string) error
	// MarkReserveRequestForIP(ip string) error
}

type client struct {
	etcd         utils.EtcdClient
	calicoDriver calico.Driver
	PoolIDV4     string
	PoolIDV6     string
}

// NewClient .
func NewClient(etcdV3 *clientv3.Client, libcalico libcalico.Interface) Client {
	return client{
		etcd:         utils.EtcdClient{Etcd: etcdV3},
		calicoDriver: calico.NewDriver(libcalico),
		PoolIDV4:     "CalicoPoolIPv4",
		PoolIDV6:     "CalicoPoolIPv6",
	}
}

func (client client) ReleaseReservedIPByTiedContainerIDIfIdle(containerID string) error {
	log.Infof("Release reserved IP by tied containerID(%s)\n", containerID)

	container := Container{ID: containerID}
	if present, err := client.etcd.GetAndDelete(&container); err != nil {
		return err
	} else if !present {
		log.Infof("the container(%s) is not exists, will do nothing\n", containerID)
		return nil
	}
	if container.Address == "" {
		log.Infof("the ip of container(%s) is empty, will do nothing\n", containerID)
		return nil
	}
	address := ReservedIPAddress{PoolID: container.PoolID, Address: container.Address}
	log.Infof("aquiring reserved address(%s)\n", address.Address)
	if present, err := client.etcd.GetAndDelete(&address); err != nil {
		log.Errorf("aquiring reserved address(%s) error", address.Address)
		return err
	} else if !present {
		log.Infof("reserved address(%s) has already been released or reallocated\n", address.Address)
		return nil
	}

	log.Infof("release ip(%s) to calico pools\n", address.Address)
	if err := client.calicoDriver.ReleaseIP(container.PoolID, container.Address); err != nil {
		log.Errorf("IP releasing error, poolID: %v, ip: %v\n", container.PoolID, container.Address)
		return err
	}

	return nil
}

func (client client) MarkReserveRequestForIP(ip string) (err error) {
	var reserved bool
	request := ReservedIPAddress{Address: ip}
	if reserved, err = client.etcd.Get(&request); reserved || err != nil {
		return
	}
	return client.etcd.Put(&request)
}

func (client client) RequestFixedIP(poolID string) (address string, err error) {
	log.Infof("Auto assigning IP from Calico pools, poolID = %s", poolID)
	var (
		ip caliconet.IP
	)
	if ip, err = client.calicoDriver.AutoAssign(poolID); err != nil {
		return "", err
	}
	return client.reserveAndFormatIPAddress(ip)
}

func (client client) reserveAndFormatIPAddress(ip caliconet.IP) (result string, err error) {
	if ip.Version() == 4 {
		// IPv4 address
		result = fmt.Sprintf("%v/%v", ip, "32")
	} else {
		// IPv6 address
		result = fmt.Sprintf("%v/%v", ip, "128")
	}
	address := fmt.Sprintf("%v", ip)
	log.Infof("[MinionsClient.reserveAndFormatIPAddress] request ip %s success, reserving...", address)
	err = client.etcd.Put(&ReservedIPAddress{Address: address})
	return
}

func (client client) MarkFixedIPForContainer(containerID string, address ReservedIPAddress) (err error) {
	return client.etcd.Put(&Container{ID: containerID, PoolID: address.PoolID, Address: address.Address})
}
