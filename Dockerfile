FROM golang:1.14.3-alpine3.11 AS BUILD

LABEL MAINTAINER="CMGS <ilskdw@gmail.com>"

# make binary
RUN apk add --no-cache ca-certificates curl make alpine-sdk linux-headers
WORKDIR /go/src/github.com/nyanpassu/minions
COPY . /go/src/github.com/nyanpassu/minions/
RUN make build && ./eru-minions --version

FROM alpine:3.11

LABEL MAINTAINER="CMGS <ilskdw@gmail.com>"

RUN mkdir /etc/eru/
COPY --from=BUILD /go/src/github.com/nyanpassu/minions/eru-minions /usr/bin/eru-minions
COPY minions.conf /usr/bin/eru-minions/