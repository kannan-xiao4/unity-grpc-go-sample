## Installation
## protoc, go
## go get google.golang.org/grpc // grpc
## go get -u github.com/golang/protobuf/protoc-gen-go // generate .pb.go plugin

# clean up old pb
rm -rf goserver/src/sample/pb/*

protoc --proto_path protocolbuffer/ protocolbuffer/* --go_out=plugins=grpc:goserver/src/sample/pb