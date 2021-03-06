package builder

var tmpl = `FROM $base AS builder

ADD . .

$path

$add-on

$run

FROM $os

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

RUN set -eux \
 && mkdir /logs \
 && ln -sf /dev/stdout /logs/out.log \
 && ln -sf /dev/stderr /logs/err.log

COPY --from=builder /app .
CMD exec ./app 1>>/logs/out.log 2>>/logs/err.log
`

var htmlTmpl = `
FROM alpine:3.4
RUN apk add --no-cache darkhttpd && mkdir -p /www
VOLUME /www
COPY . /www
EXPOSE 80
CMD ["darkhttpd", "/www"]
`

var shTmpl = `#!/bin/sh

set -eux

cd "$(dirname "$0")"

if [ -z "$GOPATH" ]; then
	export GOPATH=~/go
fi

docker build -t $projectName:$projectTag .

dangling_docker=$(docker images -f 'dangling=true' -q)
if [ -z "$dangling_docker" ]; then
    exit 1
fi

docker rmi $dangling_docker --force`

//var shTmpl = `#!/bin/sh
//
//set -eu
//
//cd "$(dirname "$0")"
//
//docker build -t $projectName:$projectTag .
//
//dangling_docker=$(docker images -f 'dangling=true' -q)
//if [ -z "$dangling_docker" ]; then
//    exit 1
//fi
//
//docker rmi $dangling_docker --force
//`
