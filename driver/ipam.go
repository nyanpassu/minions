package driver

import (
	"fmt"
	"net"

	"github.com/docker/go-plugins-helpers/ipam"
	"github.com/pkg/errors"
	caliconet "github.com/projectcalico/libcalico-go/lib/net"
	logutils "github.com/projectcalico/libnetwork-plugin/utils/log"
	"github.com/projecteru2/minions/calico"
	"github.com/projecteru2/minions/lib"
	log "github.com/sirupsen/logrus"
)

// IpamDriver .
type IpamDriver struct {
	calicoDriver calico.Driver
	ripam        ReservedIPManager

	// poolIDV4 string
	// poolIDV6 string
}

// NewIpamDriver .
func NewIpamDriver(calicoDriver calico.Driver, ripam ReservedIPManager) ipam.Ipam {
	return IpamDriver{
		calicoDriver: calicoDriver,
		ripam:        ripam,
	}
}

// GetCapabilities .
func (driver IpamDriver) GetCapabilities() (*ipam.CapabilitiesResponse, error) {
	resp := ipam.CapabilitiesResponse{}
	logutils.JSONMessage("GetCapabilities response", resp)
	return &resp, nil
}

// GetDefaultAddressSpaces .
func (driver IpamDriver) GetDefaultAddressSpaces() (*ipam.AddressSpacesResponse, error) {
	resp := &ipam.AddressSpacesResponse{
		LocalDefaultAddressSpace:  CalicoLocalAddressSpace,
		GlobalDefaultAddressSpace: CalicoGlobalAddressSpace,
	}
	logutils.JSONMessage("GetDefaultAddressSpace response", resp)
	return resp, nil
}

// RequestPool .
func (driver IpamDriver) RequestPool(request *ipam.RequestPoolRequest) (*ipam.RequestPoolResponse, error) {
	logutils.JSONMessage("RequestPool", request)

	// Calico IPAM does not allow you to request a SubPool.
	if request.SubPool != "" {
		err := errors.New(
			"Calico IPAM does not support sub pool configuration " +
				"on 'docker create network'. Calico IP Pools " +
				"should be configured first and IP assignment is " +
				"from those pre-configured pools.",
		)
		log.Errorln(err)
		return nil, err
	}

	if len(request.Options) != 0 {
		err := errors.New("Arbitrary options are not supported")
		log.Errorln(err)
		return nil, err
	}

	var (
		pool calico.Pool
		err  error
	)

	// If a pool (subnet on the CLI) is specified, it must match one of the
	// preconfigured Calico pools.
	if request.Pool != "" {
		if pool, err = driver.calicoDriver.RequestPool(request.Pool); err != nil {
			log.Errorf("[IPAMDriver::RequestPool] request calico pool error, %v", err)
			return nil, err
		}
	} else {
		pool = driver.calicoDriver.RequestDefaultPool(request.V6)
	}

	// We use static pool ID and CIDR. We don't need to signal the
	// The meta data includes a dummy gateway address. This prevents libnetwork
	// from requesting a gateway address from the pool since for a Calico
	// network our gateway is set to a special IP.
	resp := &ipam.RequestPoolResponse{
		PoolID: pool.Name,
		Pool:   pool.CIDR,
		Data:   map[string]string{"com.docker.network.gateway": pool.Gateway},
	}
	logutils.JSONMessage("RequestPool response", resp)
	return resp, nil
}

// ReleasePool .
func (driver IpamDriver) ReleasePool(request *ipam.ReleasePoolRequest) error {
	logutils.JSONMessage("ReleasePool", request)
	return nil
}

// RequestAddress .
func (driver IpamDriver) RequestAddress(request *ipam.RequestAddressRequest) (*ipam.RequestAddressResponse, error) {
	logutils.JSONMessage("RequestAddress", request)

	// Calico IPAM does not allow you to choose a gateway.
	if err := checkOptions(request.Options); err != nil {
		log.Errorf("[IpamDriver::RequestAddress] check request options failed, %v", err)
		return nil, err
	}

	var address caliconet.IP
	var err error
	if address, err = driver.requestIP(request); err != nil {
		return nil, err
	}

	// we should remove the request mark
	ip := fmt.Sprintf("%v", address)
	if _, err := driver.ripam.ConsumeRequestMarkIfPresent(lib.ReserveRequest{
		PoolID:  request.PoolID,
		Address: ip,
	}); err != nil {
		// Do not continue, or else the mark will cause some undefined behavior
		log.Errorf("[IPAM.RequestAddress] remove request mark of ip(%v) error, %v", ip, err)
		return nil, err
	} else {
		log.Infof("[IPAM.RequestAddress] removed request mark on ip(%v) allocated", ip)
	}

	resp := &ipam.RequestAddressResponse{
		// Return the IP as a CIDR.
		Address: formatIPAddress(address),
	}
	logutils.JSONMessage("RequestAddress response", resp)
	return resp, nil
}

func checkOptions(options map[string]string) error {
	// Calico IPAM does not allow you to choose a gateway.
	if options["RequestAddressType"] == "com.docker.network.gateway" {
		err := errors.New("Calico IPAM does not support specifying a gateway")
		return err
	}
	return nil
}

func (driver IpamDriver) requestIP(request *ipam.RequestAddressRequest) (caliconet.IP, error) {
	if request.Address == "" {
		return driver.calicoDriver.AutoAssign(request.PoolID)
	}
	var err error

	// specified address requested, so will try assign from reserved pool, then calico pool
	log.Println("Assigning specified IP from reserved pool first, then calico pools")

	// try to aquire ip from reserved ip pool
	var aquired bool
	if aquired, err = driver.ripam.AquireIfReserved(lib.ReservedIPAddress{
		PoolID:  request.PoolID,
		Address: request.Address,
	}); err != nil {
		return caliconet.IP{}, err
	}
	if aquired {
		return caliconet.IP{IP: net.ParseIP(request.Address)}, nil
	}
	// assign IP from calico
	return driver.calicoDriver.AssignIP(request.PoolID, request.Address)
}

func formatIPAddress(ip caliconet.IP) string {
	if ip.Version() == 4 {
		// IPv4 address
		return fmt.Sprintf("%v/%v", ip, "32")
	}
	// IPv6 address
	return fmt.Sprintf("%v/%v", ip, "128")
}

// ReleaseAddress .
func (driver IpamDriver) ReleaseAddress(request *ipam.ReleaseAddressRequest) error {
	logutils.JSONMessage("ReleaseAddress", request)

	reserved, err := driver.ripam.IsReserved(lib.ReservedIPAddress{
		PoolID:  request.PoolID,
		Address: request.Address,
	})
	if err != nil {
		log.Errorf("Get reserved ip status error, ip: %v", request.Address)
		return err
	}

	if reserved {
		log.Infof("Ip is reserved, will not release to pool, ip: %v\n", request.Address)
		return nil
	}

	return driver.calicoDriver.ReleaseIP(request.PoolID, request.Address)
}
