#!/bin/bash

unset GOROOT
export GOPATH=$(pwd)

go run challenge_three < src/challenge_three/input.data
