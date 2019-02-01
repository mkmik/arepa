#!/usr/bin/env bash

set -euo pipefail

if [[ "$1" == "test" ]]; then
    echo "Running go test"
    go build ./...
    go test -v -mod=vendor ./...
fi
