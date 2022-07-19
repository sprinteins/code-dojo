#!/bin/sh
echo ""

go get github.com/cortesi/modd/cmd/modd 
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega

cd /app && go mod download && modd
