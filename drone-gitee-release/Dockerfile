FROM alpine:3.15
LABEL Repo="https://github.com/orangedeng/drone-ossutils" Maintainer="orangedeng"
RUN apk update && apk add ca-certificates curl bash && cp /bin/bash /bin/sh
ADD create-release.sh /usr/bin/
ENTRYPOINT [ "create-release.sh" ] 