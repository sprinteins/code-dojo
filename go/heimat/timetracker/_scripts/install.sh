#!/bin/sh

cd /app

go get github.com/cortesi/modd/cmd/modd 
go get gotest.tools/gotestsum
go get golang.org/x/tools/cmd/goimports



go mod download 