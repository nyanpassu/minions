package driver

import (
	"github.com/docker/go-plugins-helpers/network"
	"github.com/projectcalico/libcalico-go/lib/clientv3"

	// "github.com/projecteru2/minions/barrel"
	calNetDriver "github.com/projecteru2/minions/driver/calico/network"
)

type NetworkDriver struct {
	calNetDriver calNetDriver.NetworkDriver
}

// NewNetworkDriver .
func NewNetworkDriver(
	client clientv3.Interface,
	// dockerCli *dockerClient.Client,
	// meta barrel.MetaInterface,
) network.Driver {
	return NetworkDriver{}
}

// GetCapabilities .
func (driver NetworkDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return driver.calNetDriver.GetCapabilities()
}

// AllocateNetwork .
func (driver NetworkDriver) AllocateNetwork(request *network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	return driver.calNetDriver.AllocateNetwork(request)
}

// FreeNetwork is used for swarm-mode support in remote plugins, which
// Calico's libnetwork-plugin doesn't currently support.
func (driver NetworkDriver) FreeNetwork(request *network.FreeNetworkRequest) error {
	return driver.calNetDriver.FreeNetwork(request)
}

// CreateNetwork .
func (driver NetworkDriver) CreateNetwork(request *network.CreateNetworkRequest) error {
	return driver.calNetDriver.CreateNetwork(request)
}

// DeleteNetwork .
func (driver NetworkDriver) DeleteNetwork(request *network.DeleteNetworkRequest) error {
	return driver.calNetDriver.DeleteNetwork(request)
}

// CreateEndpoint .
func (driver NetworkDriver) CreateEndpoint(request *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	return driver.calNetDriver.CreateEndpoint(request)
}

// DeleteEndpoint .
func (driver NetworkDriver) DeleteEndpoint(request *network.DeleteEndpointRequest) error {
	return driver.calNetDriver.DeleteEndpoint(request)
}

// EndpointInfo .
func (driver NetworkDriver) EndpointInfo(request *network.InfoRequest) (*network.InfoResponse, error) {
	return driver.calNetDriver.EndpointInfo(request)
}

// Join .
func (driver NetworkDriver) Join(request *network.JoinRequest) (*network.JoinResponse, error) {
	return driver.calNetDriver.Join(request)
}

// Leave .
func (driver NetworkDriver) Leave(request *network.LeaveRequest) error {
	return driver.calNetDriver.Leave(request)
}

// DiscoverNew .
func (driver NetworkDriver) DiscoverNew(request *network.DiscoveryNotification) error {
	return driver.calNetDriver.DiscoverNew(request)
}

// DiscoverDelete .
func (driver NetworkDriver) DiscoverDelete(request *network.DiscoveryNotification) error {
	return driver.calNetDriver.DiscoverDelete(request)
}

// ProgramExternalConnectivity .
func (driver NetworkDriver) ProgramExternalConnectivity(request *network.ProgramExternalConnectivityRequest) error {
	return driver.calNetDriver.ProgramExternalConnectivity(request)
}

// RevokeExternalConnectivity .
func (driver NetworkDriver) RevokeExternalConnectivity(request *network.RevokeExternalConnectivityRequest) error {
	return driver.calNetDriver.RevokeExternalConnectivity(request)
}
