#!/bin/sh
echo "generating protobuf codes"
rm -rf ./api/*.pb.go
protoc ./api/src/*.proto -I. --go_out=plugins=grpc:${GOPATH}/src
