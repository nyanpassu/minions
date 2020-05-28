.PHONY: deps build test binary

REPO_PATH := github.com/projecteru2/minions
REVISION := $(shell git rev-parse HEAD || unknown)
BUILTAT := $(shell date +%Y-%m-%dT%H:%M:%S)
BUILD_PATH := target
VERSION := $(shell cat VERSION)
GO_LDFLAGS ?= -s -X $(REPO_PATH)/versioninfo.REVISION=$(REVISION) \
			  -X $(REPO_PATH)/versioninfo.BUILTAT=$(BUILTAT) \
			  -X $(REPO_PATH)/versioninfo.VERSION=$(VERSION)

deps:
	go get

binary:
	@go build -ldflags "$(GO_LDFLAGS)" -a -tags netgo -installsuffix netgo -o $(BUILD_PATH)/eru-minions cmd/minions/minions.go
	@go build -ldflags "$(GO_LDFLAGS)" -a -tags netgo -installsuffix netgo -o $(BUILD_PATH)/minionsctl cmd/minionsctl/minionsctl.go

debug-binary:
	@go build -ldflags "$(GO_LDFLAGS)" -gcflags "-N -l" -a -tags netgo -installsuffix netgo -o $(BUILD_PATH)/eru-minions cmd/minions/minions.go
	@go build -ldflags "$(GO_LDFLAGS)" -gcflags "-N -l" -a -tags netgo -installsuffix netgo -o $(BUILD_PATH)/minionsctl cmd/minionsctl/minionsctl.go

build: deps binary
debug-build: deps debug-binary

docker_build:
	docker run --rm \
		-v $(CURDIR):/go/src/github.com/nyanpassu/minions:rw \
		-v $(HOME)/.glide:/home/user/.glide:rw \
		-e LOCAL_USER_ID=$(LOCAL_USER_ID) \
		$(GO_BUILD_CONTAINER) /bin/sh -c ' \
			cd /go/src/github.com/projectcalico/libnetwork-plugin && \
			glide install -strip-vendor' 

test: deps
	# fix mock docker client bug, see https://github.com/moby/moby/pull/34383 [docker 17.05.0-ce]
	sed -i.bak "143s/\*http.Transport/http.RoundTripper/" ./vendor/github.com/docker/docker/client/client.go
	go vet `go list ./... | grep -v '/vendor/'`
	go test -cover -v `glide nv`
