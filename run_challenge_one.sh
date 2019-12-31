#!/bin/bash

unset GOROOT
export GOPATH=$(pwd)

go run challenge01 < src/challenge01/module_masses.data
