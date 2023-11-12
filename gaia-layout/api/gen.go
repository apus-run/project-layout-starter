package api

//go:generate protoc -I . -I ../third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./helloworld/v1/greeter.proto
//go:generate protoc -I . -I ../third_party --go_out=paths=source_relative:. --go-gin_out=paths=source_relative:. ./helloworld/v1/greeter.proto
//go:generate protoc -I . -I ../third_party --go_out=paths=source_relative:. --openapi_out=paths=source_relative:. ./helloworld/v1/greeter.proto

//go:generate protoc -I . -I ../third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./user/v1/user.proto
//go:generate protoc -I . -I ../third_party --go_out=paths=source_relative:. --go-gin_out=paths=source_relative:. ./user/v1/user.proto
//go:generate protoc -I . -I ../third_party --go_out=paths=source_relative:. --openapi_out=paths=source_relative:. ./user/v1/user.proto
