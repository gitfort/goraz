#!/bin/sh
echo "installing dependencies"
go get ./...
go mod tidy
