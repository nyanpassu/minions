package driver

import (
	"fmt"
	"os"
	"strings"
	"time"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/pkg/errors"
	caliconet "github.com/projectcalico/libcalico-go/lib/net"
	log "github.com/sirupsen/logrus"
)

func formatIPAddress(ip caliconet.IP) string {
	if ip.Version() == 4 {
		// IPv4 address
		return fmt.Sprintf("%v/%v", ip, "32")
	}
	// IPv6 address
	return fmt.Sprintf("%v/%v", ip, "128")
}

func checkOptions(options map[string]string) error {
	// Calico IPAM does not allow you to choose a gateway.
	if options["RequestAddressType"] == "com.docker.network.gateway" {
		err := errors.New("Calico IPAM does not support specifying a gateway")
		return err
	}
	return nil
}

func containerHasFixedIPLabel(container dockerTypes.Container) bool {
	value, hasFixedIPLabel := container.Labels[fixedIPLabel]
	return hasFixedIPLabel && strings.ToLower(value) != "false" && value != "0"
}

// Returns the label poll timeout. Default is returned unless an environment
// key is set to a valid time.Duration.
func getLabelPollTimeout() time.Duration {
	// 5 seconds should be more than enough for this plugin to get the
	// container labels. More info in func populateWorkloadEndpointWithLabels
	defaultTimeout := 5 * time.Second

	timeoutVal := os.Getenv(LABEL_POLL_TIMEOUT_ENVKEY)
	if timeoutVal == "" {
		return defaultTimeout
	}

	labelPollTimeout, err := time.ParseDuration(timeoutVal)
	if err != nil {
		err = errors.Wrapf(err, "Label poll timeout specified via env key %s is invalid, using default %s",
			LABEL_POLL_TIMEOUT_ENVKEY, defaultTimeout)
		log.Warningln(err)
		return defaultTimeout
	}
	log.Infof("Using custom label poll timeout: %s", labelPollTimeout)
	return labelPollTimeout
}
