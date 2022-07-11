package main

import tag "github.com/go-funcards/tag-service/cmd"

//go:generate protoc -I proto --go_out=./proto/v1 --go-grpc_out=./proto/v1 proto/v1/tag.proto

func main() {
	tag.Execute()
}
