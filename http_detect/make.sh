#!/bin/bash

BASE=$(pwd)

cd $GOPATH/src/github.com/toomore/Pensieve-go/http_detect
go get -v ./...

cd $GOPATH/bin
cp ./http_detect $BASE

cd $BASE
docker build -t toomore/http_detect .

rm -rf ./http_detect
