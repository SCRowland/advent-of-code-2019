#!/bin/bash

unset GOROOT
export GOPATH=$(pwd)

go run challenge03 < src/challenge03/input.data
