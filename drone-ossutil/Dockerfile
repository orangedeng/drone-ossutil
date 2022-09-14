FROM alpine:3.15
LABEL Repo="https://github.com/orangedeng/drone-ossutils" Maintainer="orangedeng"
RUN apk update && apk add ca-certificates
ARG OSSUTIL_VERSION=v1.7.13
ARG CONFIGD_VERSION=v0.16.0

RUN mkdir -p /usr/local/bin && \
    wget http://gosspublic.alicdn.com/ossutil/${OSSUTIL_VERSION##v}/ossutil64 -O /usr/local/bin/ossutil && chmod +x /usr/local/bin/ossutil && \
    wget https://github.com/kelseyhightower/confd/releases/download/${CONFIGD_VERSION}/confd-${CONFIGD_VERSION##v}-linux-amd64 -O /usr/local/bin/confd && chmod +x /usr/local/bin/confd && mkdir -p /etc/confd/conf.d /etc/confd/templates
ADD ossconfig.tmpl /etc/confd/templates/
ADD ossconfig.toml /etc/confd/conf.d/
ADD entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
