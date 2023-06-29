FROM debian:bullseye
ENV BUILDX_VERSION=v0.9.1
ARG IMAGE_MIRROR_VERSION=47e65fe50acd8b5b5dabc350eecee93997ff5787

RUN apt-get update \
    && apt-get install -y \
    skopeo \
    docker \
    jq \
    bash \
    curl \
    && cp -r /bin/bash /bin/sh \
    && mkdir -p /root/.docker/cli-plugins/ && curl -sLo /root/.docker/cli-plugins/docker-buildx https://github.com/docker/buildx/releases/download/${BUILDX_VERSION}/buildx-${BUILDX_VERSION}.linux-amd64 && chmod a+x /root/.docker/cli-plugins/* \
    && mkdir -p /image-mirror
WORKDIR /image-mirror
# Get scripts from github.com/rancher/image-mirror
RUN for script in compare.sh image-mirror.sh sort-images-list.sh; do \
    curl -sSL https://raw.githubusercontent.com/rancher/image-mirror/${IMAGE_MIRROR_VERSION}/$script > /image-mirror/$script; \
    chmod +x /image-mirror/$script; \
    done;
# Add scripts
COPY drone-entrypoint.sh /image-mirror/

ENTRYPOINT ["/image-mirror/drone-entrypoint.sh"]
