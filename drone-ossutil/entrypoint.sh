#!/bin/sh

set -e

if [ -z "${ENDPOINT}" ] || [ -z "${ACCESS_KEY_ID}" ] || [ -z "${ACCESS_KEY_SECRET}" ]; then
    echo "endpoint, access key id or access key secret not set."
    exit 1
fi

confd -onetime -backend env

if [ "$?" -ne "0" ];then
    exit $?
fi

exec "$@"