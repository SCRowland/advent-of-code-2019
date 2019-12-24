#!/bin/bash

unset GOROOT
export GOPATH=$(pwd)

go run challenge_five < src/challenge_five/machine_code.data
