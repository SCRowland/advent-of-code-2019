#!/bin/bash

unset GOROOT
export GOPATH=$(pwd)

go run challenge02 < src/challenge02/machine_code.data
