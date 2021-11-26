#!/bin/bash

set -e

python3 -m pip install grpcio
python3 -m pip install grpcio-tools

GEN_PATH="."
GO_LIB_PATH=$(go env GOPATH)/src
GOPATH=$(go env GOPATH)

mkdir -p proto/google/api 
curl -o proto/google/api/annotations.proto "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto"
curl -o proto/google/api/http.proto "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto"

mkdir -p proto/google/protobuf
curl -o proto/google/protobuf/descriptor.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/descriptor.proto"
curl -o proto/google/protobuf/empty.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/empty.proto"
curl -o proto/google/protobuf/struct.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/struct.proto"
curl -o proto/google/protobuf/timestamp.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/timestamp.proto"
curl -o proto/google/protobuf/any.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/any.proto"

protoc \
        -I proto \
        -I $GO_LIB_PATH/include \
        --go_out=$GEN_PATH      --go_opt=paths=source_relative \
        --go-grpc_out=$GEN_PATH --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=logtostderr=true,paths=source_relative:$GEN_PATH \
        --openapiv2_out=logtostderr=true:. \
        proto/groups.proto

python3 -m grpc_tools.protoc \
        -I proto \
        -I $GOPATH/src/include \
        --python_out=$GEN_PATH/groups \
        $GOPATH/src/include/protoc-gen-openapiv2/options/annotations.proto

python3 -m grpc_tools.protoc \
        -I proto \
        -I $GOPATH/src/include \
        --python_out=$GEN_PATH/groups \
        $GOPATH/src/include/protoc-gen-openapiv2/options/openapiv2.proto

python3 -m grpc_tools.protoc \
        -I proto \
        -I $GOPATH/src/include \
        --python_out=$GEN_PATH/groups \
        --grpc_python_out=$GEN_PATH/groups \
        proto/groups.proto

echo "inject tag"
protoc-go-inject-tag -input=$GEN_PATH/groups.pb.go

echo "update and tidy"
go mod tidy
go get -u ./...
