#!/bin/bash

unset GOROOT
export GOPATH=$(pwd)

go run challenge_two < src/challenge_two/machine_code.data
