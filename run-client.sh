#!/usr/bin/env bash

# Runs amolecoin in desktop client configuration

set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "amolecoin binary dir:" "$DIR"
pushd "$DIR" >/dev/null

COMMIT=$(git rev-parse HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
GOLDFLAGS="${GOLDFLAGS} -X main.Commit=${COMMIT} -X main.Branch=${BRANCH}"

GORUNFLAGS=${GORUNFLAGS:-}

go run -ldflags "${GOLDFLAGS}" $GORUNFLAGS ./cmd/amolecoin/... \
    -gui-dir="${DIR}/src/gui/static/" \
    -launch-browser=true \
    -enable-all-api-sets=true \
    -enable-gui=true \
    -log-level=debug \
    $@

popd >/dev/null