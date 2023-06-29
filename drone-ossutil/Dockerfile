FROM registry.suse.com/bci/bci-base:15.5
LABEL Repo="https://github.com/orangedeng/drone-ossutils" Maintainer="orangedeng"
ARG TARGETARCH
RUN zypper install -y wget unzip
ARG OSSUTIL_VERSION=v1.7.16
ARG CONFIGD_VERSION=v0.16.0

RUN mkdir -p /usr/local/bin && \
    wget http://gosspublic.alicdn.com/ossutil/${OSSUTIL_VERSION##v}/ossutil-${OSSUTIL_VERSION}-linux-${TARGETARCH}.zip -O /tmp/ossutil.zip && \ 
    unzip -a /tmp/ossutil.zip -d /tmp && cd /tmp/ossutil-* && chmod 0755 ./ossutil && chown root:root ./ossutil && mv ./ossutil /usr/local/bin/ossutil && \
    wget https://github.com/kelseyhightower/confd/releases/download/${CONFIGD_VERSION}/confd-${CONFIGD_VERSION##v}-linux-${TARGETARCH} -O /usr/local/bin/confd && chmod +x /usr/local/bin/confd && mkdir -p /etc/confd/conf.d /etc/confd/templates
ADD ossconfig.tmpl /etc/confd/templates/
ADD ossconfig.toml /etc/confd/conf.d/
ADD entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
