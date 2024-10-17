#!/bin/sh
SHELL_COMMAND="./demo.sh" USERNAME=user PASSWORD=pass1234 go run `ls *.go | grep -v _test.go`