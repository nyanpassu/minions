module github.com/projecteru2/minions

go 1.14

require (
	cloud.google.com/go v0.1.1-0.20160913182117-3b1ae45394a2
	github.com/Azure/go-autorest v8.0.0+incompatible
	github.com/Microsoft/go-winio v0.4.5
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/PuerkitoBio/purell v1.0.0
	github.com/PuerkitoBio/urlesc v0.0.0-20160726150825-5bd2802263f2
	github.com/codegangsta/cli v1.20.0
	github.com/coreos/etcd v3.3.8+incompatible
	github.com/coreos/go-systemd v0.0.0-20161114122254-48702e0da86b
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.0.0+incompatible
	github.com/docker/distribution v2.6.0-rc.1.0.20171207180435-f4118485915a+incompatible
	github.com/docker/docker v17.12.0-ce-rc1.0.20180616010903-de0abf4315fd+incompatible
	github.com/docker/go-connections v0.3.0
	github.com/docker/go-plugins-helpers v0.0.0-20170919092928-bd8c600f0cdd
	github.com/docker/go-units v0.3.2
	github.com/emicklei/go-restful v1.1.4-0.20151126145626-777bb3f19bca
	github.com/emicklei/go-restful-swagger12 v0.0.0-20170208215640-dcef7f557305
	github.com/ghodss/yaml v0.0.0-20150909031657-73d445a93680
	github.com/go-openapi/jsonpointer v0.0.0-20160704185906-46af16f9f7b1
	github.com/go-openapi/jsonreference v0.0.0-20160704190145-13c6e3589ad9
	github.com/go-openapi/spec v0.0.0-20160808142527-6aced65f8501
	github.com/go-openapi/swag v0.0.0-20160704191624-1d0bd113de87
	github.com/gogo/protobuf v1.1.1
	github.com/golang/glog v0.0.0-20141105023935-44145f04b68c
	github.com/golang/protobuf v1.3.2
	github.com/google/btree v0.0.0-20161005200959-925471ac9e21
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/google/gofuzz v0.0.0-20161122191042-44d81051d367
	github.com/googleapis/gnostic v0.0.0-20170426233943-68f4ded48ba9
	github.com/gophercloud/gophercloud v0.0.0-20170831144856-2bf16b94fdd9
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/gregjones/httpcache v0.0.0-20170728041850-787624de3eb7
	github.com/hashicorp/golang-lru v0.0.0-20160207214719-a0d98a5f2880
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	github.com/imdario/mergo v0.0.0-20141206190957-6633656539c1
	github.com/json-iterator/go v1.1.6
	github.com/juju/ratelimit v0.0.0-20170523012141-5b9ff8664717
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/mailru/easyjson v0.0.0-20160728113105-d5b7844b561a
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/opencontainers/image-spec v1.0.1-0.20180411145040-e562b0440392
	github.com/pborman/uuid v0.0.0-20150603214016-ca53cad383ca
	github.com/peterbourgon/diskv v2.0.1+incompatible
	github.com/pkg/errors v0.8.1
	github.com/projectcalico/go-json v0.0.0-20161128004156-6219dc7339ba
	github.com/projectcalico/go-yaml v0.0.0-20161201183616-955bc3e451ef
	github.com/projectcalico/go-yaml-wrapper v0.0.0-20161127220527-598e54215bee
	github.com/projectcalico/libcalico-go v2.0.0-alpha1.0.20180615230155-efdf8fede805+incompatible
	github.com/projectcalico/libnetwork-plugin v1.1.3-0.20180524185918-f42c4fce3cdb
	github.com/prometheus/common v0.10.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v0.0.0-20170130214245-9ff6c6923cff
	github.com/vishvananda/netlink v0.0.0-20171214173445-54ad9e3a4cbb
	github.com/vishvananda/netns v0.0.0-20171111001504-be1fbeda1936
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980
	golang.org/x/oauth2 v0.0.0-20170412232759-a6bd8cefa181
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894
	golang.org/x/text v0.3.0
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/appengine v0.0.0-20160823001527-4f7eeb5305a4
	google.golang.org/genproto v0.0.0-20170731182057-09f6ed296fc6
	google.golang.org/grpc v1.7.5
	gopkg.in/go-playground/validator.v8 v8.18.1
	gopkg.in/inf.v0 v0.9.0
	gopkg.in/yaml.v2 v2.2.4
	gotest.tools v2.2.0+incompatible // indirect
	k8s.io/api v0.0.0-20180510182548-a315a049e7a9
	k8s.io/apimachinery v0.0.0-20180510182146-40eaf68ee188
	k8s.io/client-go v0.0.0-20170922112243-82aa063804cf
	k8s.io/kube-openapi v0.0.0-20180509233829-0c329704159e
)
