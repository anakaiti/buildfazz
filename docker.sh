#!/bin/sh

set -eu

cd "$(dirname "$0")"

docker build -t a1:a2 .

dangling_docker=$(docker images -f 'dangling=true' -q)
if [ -z "$dangling_docker" ]; then
    exit 1
fi

docker rmi $dangling_docker --force