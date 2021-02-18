#!/bin/bash

OUTPUT="server"
VERSION="1.0.0"
OS=("$(go env GOOS)")
ARCH=("$(go env GOARCH)")

echo "==> Build for ${OS}/${ARCH} with bin name: ${OUTPUT}, version: ${VERSION}"
GOOS="${OS}" GOARCH="${ARCH}" go build -o ./"${OUTPUT}" cmd/app/server.go
if [ $? -ne 0 ]; then
    exit 1
fi
echo "Build done: ${OUTPUT}."