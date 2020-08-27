module github.com/projecteru2/minions

go 1.14

require (
	github.com/Azure/go-autorest/autorest v0.11.4 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.2 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20180616010903-de0abf4315fd+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-plugins-helpers v0.0.0-20200102110956-c9a8a2d92ccc
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/gophercloud/gophercloud v0.12.0 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/projectcalico/go-json v0.0.0-20161128004156-6219dc7339ba // indirect
	github.com/projectcalico/go-yaml-wrapper v0.0.0-20191112210931-090425220c54 // indirect
	github.com/projectcalico/libcalico-go v3.4.0-0.dev+incompatible
	github.com/projectcalico/libnetwork-plugin v1.1.3-0.20180524185918-f42c4fce3cdb
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.4.0
	github.com/urfave/cli/v2 v2.2.0
	github.com/vishvananda/netlink v1.1.0
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200824191128-ae9734ed278b
	go.uber.org/automaxprocs v1.3.0
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	k8s.io/api v0.15.12
	k8s.io/apimachinery v0.15.12
	k8s.io/client-go v0.15.12
	k8s.io/kube-openapi v0.0.0-20200811211545-daf3cbb84823 // indirect
	k8s.io/utils v0.0.0-20200821003339-5e75c0163111 // indirect
)

replace github.com/coreos/etcd => go.etcd.io/etcd v3.3.25+incompatible
