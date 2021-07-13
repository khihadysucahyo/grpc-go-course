protoc --go_out=plugins=grpc:. ./greet/greetpb/greet.proto
go run ./greet/greet_server/server.go