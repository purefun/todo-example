package commands

// install

// go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
// go install github.com/asynkron/protoactor-go/protobuf/protoc-gen-gograinv2@latest

//go:generate protoc --go_out=. --go_opt=paths=source_relative --proto_path=. commands.proto
//go:generate protoc -I=. --gograinv2_out=. commands.proto
