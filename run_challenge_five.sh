#!/bin/bash

unset GOROOT
export GOPATH=$(pwd)

go run challenge05 < src/challenge05/machine_code.data
