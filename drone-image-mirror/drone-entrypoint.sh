#!/bin/bash

set -e

BASEDIR=$(dirname $0)

if [ -n "${DEBUG}" ]; then
    set -x
fi

if [ -e /run/drone/env ]; then
    source /run/drone/env
fi

LIST_FILE=${PLUGIN_IMAGE_LIST:-image-list}
DEST_REGISTRY_OVERRIDE=${PLUGIN_TARGET_REGISTRY:-docker.io}

if [ -z "${DEST_REGISTRY_OVERRIDE}" ]; then 
    echo "target registry is required";
    exit 1;
fi

DOCKER_USERNAME=${PLUGIN_DOCKER_USERNAME}
DOCKER_PASSWORD=${PLUGIN_DOCKER_PASSWORD}

if [ -n "${DOCKER_USERNAME}" -a -n "${DOCKER_PASSWORD}" ]; then 
    echo "Logging in to ${DOCKER_REGISTRY:-docker.io} as ${DOCKER_USERNAME}"
    docker login ${DEST_REGISTRY_OVERRIDE} --username=${DOCKER_USERNAME} --password-stdin <<< ${DOCKER_PASSWORD}
fi

exec cat ${LIST_FILE} | DEST_REGISTRY_OVERRIDE=${DEST_REGISTRY_OVERRIDE} ${BASEDIR}/image-mirror.sh
