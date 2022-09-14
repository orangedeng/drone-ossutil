#!/bin/sh

set -e

if [ -e /run/drone/env ]; then
    source /run/drone/env
fi

DRONE_TAG=${DRONE_TAG}
DRONE_COMMIT=${DRONE_COMMIT}
PRERELEASE=${PLUGIN_PRERELEASE:-${GITEE_RELEASE_PRERELEASE}}
GITEE_TOKEN=${PLUGIN_ACCESS_TOKEN:-${GITEE_RELEASE_ACCESS_TOKEN}}
TITLE=${PLUGIN_TITLE:-${GITEE_RELEASE_TITLE}}
NOTE=${PLUGIN_NOTE:-${GITEE_RELEASE_NOTE}}
DRONE_REPO=${DRONE_REPO}


data="{\"access_token\":\"${GITEE_TOKEN}\",\"tag_name\":\"${DRONE_TAG}\",\"name\":\"${TITLE}\",\"body\":\"${NOTE}\",\"prerelease\":\"${PRERELEASE}\",\"target_commitish\":\"${DRONE_COMMIT}\"}"
url="https://gitee.com/api/v5/repos/${DRONE_REPO}/releases"

curl -v -X POST --header 'Content-Type: application/json;charset=UTF-8' $url -d "${data}"

