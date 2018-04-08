#!/usr/bin/env bash

# Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
# Use of this document is governed by a license found in the LICENSE document.

source "$(git rev-parse --show-toplevel)/_scripts/common.bash"

rm -f !(_vendor)/**/gen_*.go

$go install myitcv.io/immutable/cmd/immutableGen
$go install myitcv.io/immutable/cmd/immutableVet

pushd cmd/immutableVet/_testFiles
$go generate
popd

$go generate ./...

# for immutable this step needs to be before the tests (because some of the
# tests rely on the myitcv.io/immutable package to have been installed)

# we can remove this once we resolve https://github.com/golang/go/issues/24661
$go install ./...

$go test ./...


immutableVet myitcv.io/immutable/example
