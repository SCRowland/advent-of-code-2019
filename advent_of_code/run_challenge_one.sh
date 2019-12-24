#!/bin/bash

unset GOROOT
export GOPATH=$(pwd)

go run challenge_one < src/challenge_one/module_masses.data
