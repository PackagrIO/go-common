#!/usr/bin/env bash

# because packagr/go-common depends on compiled binaries (even for testing) we'll be building the test binaries first in the "build container" and then
# executing them in a "runtime container" to get coverage/profiling data.
#
# this script generates the test binaries in the "build container"

set -e
go mod vendor
mkdir -p vendor/gopkg.in/libgit2/git2go.v25/vendor/libgit2/build/
cp /usr/local/linux/lib/pkgconfig/libgit2.pc vendor/gopkg.in/libgit2/git2go.v25/vendor/libgit2/build/libgit2.pc
go test -mod vendor -race -tags="static" ./...
