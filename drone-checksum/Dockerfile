FROM alpine:3.15 as runner
FROM registry.suse.com/bci/golang:1.17 as builder
ENV PKG=github.com/orangedeng/drone-tools
ENV GOPROXY=https://goproxy.cn,direct
ARG VERSION=v0.0.1
COPY . /go/src/${PKG}
WORKDIR /go/src/${PKG}
RUN CGO_ENABLED=0 go build -a -ldflags "-w -extldflags \"-static\" -X main.Version=${VERSION}" \
    -tags netgo \
    -o drone-checksum/drone-checksum drone-checksum/main.go
FROM runner
LABEL Repo="https://github.com/orangedeng/drone-ossutils" Maintainer="orangedeng"
RUN apk update && apk add ca-certificates curl bash && cp /bin/bash /bin/sh
COPY --from=builder /go/src/github.com/orangedeng/drone-tools/drone-checksum/drone-checksum /usr/bin/
ENTRYPOINT [ "drone-checksum" ]