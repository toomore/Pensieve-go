#!/usr/bin/env bash
go build
mv ./http_detect ./http_detect_darwin_amd64
GOOS=linux GOARCH=amd64 go build
mv ./http_detect ./http_detect_linux_amd64
GOOS=linux GOARCH=arm go build
mv ./http_detect ./http_detect_linux_arm
