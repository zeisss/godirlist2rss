#!/bin/bash

set -eu

docker run -ti -e GOARCH=arm -e GOARM=6 -v $(pwd):/go golang:1.5.3 go build .
scp go pi@phobos.localnet:gowalker
