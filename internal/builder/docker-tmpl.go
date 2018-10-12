package builder

var tmpl = `FROM $base AS builder

ADD . .

$path

$add-on

$run

FROM $os

RUN set -eux \
 && mkdir /logs \
 && ln -sf /dev/stdout /logs/out.txt \
 && ln -sf /dev/stderr /logs/err.txt

COPY --from=builder /app .
ENTRYPOINT ["./app"]
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
