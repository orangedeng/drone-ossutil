FROM alpine:3.15
LABEL Repo="https://github.com/orangedeng/drone-ossutils" Maintainer="orangedeng"
RUN apk update && apk add ca-certificates curl
ADD create-release.sh /use/local/bin/
CMD [ "create-release.sh " ] 